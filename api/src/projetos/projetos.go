package projetos

import (
	"database/sql"
	"errors"
	"fmt"
	"go-index-projetos/api/src/pecas"
	"log"
)

type Projeto struct {
	ID           int
	Nome         string
	Tipo         string
	Subtipo      string
	LinkCircuito string
}

func CadastrarProjeto(db *sql.DB) {
	// Cadastrar dados que vao para a tabela projetos
	var p Projeto
	var idPeca int
	fmt.Print("Digite o nome do projeto: ")
	fmt.Scanln(&p.Nome)

	fmt.Print("Digite o tipo do projeto: ")
	fmt.Scanln(&p.Tipo)

	fmt.Print("Digite o subtipo do projeto: ")
	fmt.Scanln(&p.Subtipo)

	fmt.Print("Digite o link do circuito: ")
	fmt.Scanln(&p.LinkCircuito)

	_, err := db.Exec("INSERT INTO projetos (nome, tipo, subtipo, link_circuito) VALUES (?, ?, ?, ?)", p.Nome, p.Tipo, p.Subtipo, p.LinkCircuito)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow("SELECT idproj FROM projetos WHERE nome = ?", p.Nome).Scan(&p.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Loop para receber as peças do projeto, parando para cadastrar uma peça quando o mesmo ainda nao existir na tabela pecas, e conectando o id do projeto com o id do peça na tabela proj_pe
	entradaUsuario := ""
	for i := 1; entradaUsuario != "sair"; i++ {
		var quant_nec int
		fmt.Print("Digite o valor da ", i, "a peça (ou 'sair' caso não tenha mais peças):")
		fmt.Scanln(&entradaUsuario)
		if entradaUsuario != "sair" {
			err := db.QueryRow("SELECT idpe FROM pecas WHERE valor = ?", entradaUsuario).Scan(&idPeca)
			// Verifica se essa peça já existe na tabela ou se precisa cadastrar
			if errors.Is(err, sql.ErrNoRows) {
				idPeca = pecas.CadastrarPecaDeProjeto(db, entradaUsuario)
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Print("Digite a quantidade necessaria para fazer o projeto: ")
			fmt.Scanln(&quant_nec)
			// Conecta o ID do projeto com os IDs de todas as peças que são usadas no projeto
			_, err = db.Exec("INSERT INTO proj_pe (idproj, idpe, quant_pecas) VALUES (?, ?, ?)", p.ID, idPeca, quant_nec)
		}
	}
}

func BuscarProjeto(db *sql.DB) {
	var termoBusca, colunaBusca string
	var escolha, escolhaBusca int
	var p Projeto

	for escolha != 4 {
		fmt.Println("Buscar por: \n1 - Nome\n2 - Tipo\n3 - Link\n4 - Voltar")
		fmt.Scan(&escolha)
		if escolha == 4 {
			break
		}

		switch escolha {
		case 1:
			colunaBusca = "nome"
			fmt.Println("Digite o nome do projeto: ")
			fmt.Scan(&termoBusca)
		case 2:
			colunaBusca = "tipo"
			fmt.Println("Digite o tipo de projeto: ")
			fmt.Scan(&termoBusca)
		case 3:
			colunaBusca = "link"
			fmt.Println("Digite o link do projeto: ")
			fmt.Scan(&termoBusca)
		case 4:
			break
		default:
			fmt.Println("...")
			return
		}

		query := fmt.Sprintf("SELECT idproj, nome, tipo, subtipo, link_circuito FROM projetos WHERE %s = ?", colunaBusca)
		rows, err := db.Query(query, termoBusca)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var listaProjetos []Projeto
		fmt.Print("\n---Projetos Encontrados---\n")
		contador := 1
		for rows.Next() {
			err := rows.Scan(&p.ID, &p.Nome, &p.Tipo, p.Subtipo, &p.LinkCircuito)
			if err != nil {
				log.Fatal(err)
			}
			listaProjetos = append(listaProjetos, p)
			fmt.Printf("%d - %s\n", contador, p.Nome)
			contador++
		}

		if len(listaProjetos) == 0 {
			fmt.Println("Nenhum projeto encontrado...")
			break
		}

		fmt.Println("Digite o número do projeto que você deseja: ")
		fmt.Scan(&escolhaBusca)

		if escolhaBusca < 1 || escolhaBusca > len(listaProjetos) {
			fmt.Println("Número inválido...")
			break
		}

		projetoSelecionado := listaProjetos[escolhaBusca-1]
		idproj := projetoSelecionado.ID

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Nome: %s\nTipo: %s\nSubtipo: %s\nLink do circuito: %s\n",
			projetoSelecionado.Nome, projetoSelecionado.Tipo, projetoSelecionado.Subtipo, projetoSelecionado.LinkCircuito)

		for escolha != 4 {
			fmt.Println("\nVocê deseja:\n1 - Ver peças usadas nesse projeto\n2 - Verificar se você tem todas as peças\n3 - Voltar")
			fmt.Scan(&escolha)
			switch escolha {
			case 1:
				verificarPecasProjeto(db, projetoSelecionado)
			case 2:
				if compararQuantEstoque(db, idproj) {
					fmt.Println("Você tem todas as peças necessárias para esse projeto! Deseja dar baixa no estoque e fazer o projeto?\n1 - Sim\n2 - Não")
					var escolhaBaixa int
					fmt.Scan(&escolhaBaixa)
					switch escolhaBaixa {
					case 1:
						baixaQuantEstoque(db, idproj)
						fmt.Print("Pronto!")
					case 2:
						break
					}
				} else {
					fmt.Println("Você não tem todas as peças necessárias para esse projeto.")
				}
			case 3:
				escolha = 4
			default:
				fmt.Println("Opção inválida.")
			}
		}
	}

}

func compararQuantEstoque(db *sql.DB, idproj int) bool {
	queryEstoque :=
		`SELECT p.quant_estoque, pp.quant_pecas 
				FROM proj_pe pp 
				JOIN pecas p 
				ON pp.idpe = p.idpe 
				WHERE pp.idproj = ?`
	rowsEstoque, err := db.Query(queryEstoque, idproj)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsEstoque.Close()
	temTodasPecas := true
	for rowsEstoque.Next() {
		var quantEstoque, quantNecessaria int
		err := rowsEstoque.Scan(&quantEstoque, &quantNecessaria)
		if err != nil {
			log.Fatal(err)
		}
		if quantEstoque < quantNecessaria {
			temTodasPecas = false
		}
	}
	return temTodasPecas
}

func baixaQuantEstoque(db *sql.DB, idproj int) {
	queryEstoque :=
		`SELECT pp.idpe, p.quant_estoque, pp.quant_pecas 
		FROM proj_pe pp 
		JOIN pecas p ON pp.idpe = p.idpe 
		WHERE pp.idproj = ?`
	rowsEstoque, err := db.Query(queryEstoque, idproj)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsEstoque.Close()
	for rowsEstoque.Next() {
		var idpe, quantEstoque, quantNecessaria int
		err := rowsEstoque.Scan(&idpe, &quantEstoque, &quantNecessaria)
		if err != nil {
			log.Fatal(err)
		}
		quantAtualizada := quantEstoque - quantNecessaria
		_, err = db.Exec("UPDATE pecas SET quant_estoque = ? WHERE idpe = ?", quantAtualizada, idpe)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func verificarPecasProjeto(db *sql.DB, p Projeto) {
	queryPecas :=
		`SELECT p.tipo, p.valor, p.detalhe, p.voltagem, p.quant_estoque, pp.quant_pecas 
		FROM proj_pe pp 
		JOIN pecas p 
		ON pp.idpe = p.idpe 
		WHERE pp.idproj = ?`
	rowsPecas, err := db.Query(queryPecas, p.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsPecas.Close()
	fmt.Printf("\n--- Peças usadas no %s ---\n", p.Nome)
	for rowsPecas.Next() {
		var tipoPeca, valorPeca, detalhePeca, voltagemPeca string
		var quantEstoque, quantNecessaria int
		err := rowsPecas.Scan(&tipoPeca, &valorPeca, &detalhePeca, &voltagemPeca, &quantEstoque, &quantNecessaria)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Tipo: %s | Valor: %s | Detalhe: %s | Voltagem: %s | Quantidade em estoque: %d | Quantidade necessária: %d\n", tipoPeca, valorPeca, detalhePeca, voltagemPeca, quantEstoque, quantNecessaria)
	}
	fmt.Print("-----------------------------------------------------\n")
}
