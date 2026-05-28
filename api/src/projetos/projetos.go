package projetos

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func CadastrarProjeto(db *sql.DB) {
	// Cadastrar dados que vao para a tabela projetos
	var nome, tipo, subtipo, linkCircuito string
	var idProj, idPeca int
	fmt.Print("Digite o nome do projeto: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite o tipo do projeto: ")
	fmt.Scanln(&tipo)

	fmt.Print("Digite o subtipo do projeto: ")
	fmt.Scanln(&subtipo)

	fmt.Print("Digite o link do circuito: ")
	fmt.Scanln(&linkCircuito)

	_, err := db.Exec("INSERT INTO projetos (nome, tipo, subtipo, link_circuito) VALUES (?, ?, ?, ?)", nome, tipo, subtipo, linkCircuito)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow("SELECT idproj FROM projetos WHERE nome = ?", nome).Scan(&idProj)
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
				var tipo, detalhe, voltagem string
				var quant_estoque int
				fmt.Println("Essa peça ainda não existe!\n.\n.\n.")
				fmt.Print("Digite o tipo da peça nova: ")
				fmt.Scanln(&tipo)
				fmt.Print("Digite o detalhe da peça nova: ")
				fmt.Scanln(&detalhe)
				fmt.Print("Digite a voltagem da peça nova: ")
				fmt.Scanln(&voltagem)
				fmt.Print("Digite quantidade no estoque da peça nova: ")
				fmt.Scanln(&quant_estoque)
				fmt.Println("Cadastrando peça nova...\n.\n.\n.")
				_, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", tipo, entradaUsuario, detalhe, voltagem, quant_estoque)
				if err != nil {
					log.Fatal(err)
				}
				err = db.QueryRow("SELECT LAST_INSERT_ID() FROM pecas").Scan(&idPeca)
				if err != nil {
					log.Fatal(err)
				}
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Print("Digite a quantidade necessaria para fazer o projeto: ")
			fmt.Scanln(&quant_nec)
			// Conecta o ID do projeto com os IDs de todas as peças que são usadas no projeto
			_, err = db.Exec("INSERT INTO proj_pe (idproj, idpe, quant_pecas) VALUES (?, ?, ?)", idProj, idPeca, quant_nec)
		}
	}
}

func BuscarProjeto(db *sql.DB) {
	var titulo, variavel, nome, tipo, subtipo, linkCircuito string
	var id, escolha int
	for escolha != 4 {
		fmt.Println("Buscar por: \n1 - nome\n2 - tipo\n3 - link\n4 - voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			variavel = "nome"
			fmt.Println("Digite o nome do projeto: ")
			fmt.Scan(&titulo)
		case 2:
			variavel = "tipo"
			fmt.Println("Digite o tipo de projeto: ")
			fmt.Scan(&titulo)
		case 3:
			variavel = "link"
			fmt.Println("Digite o link do projeto: ")
			fmt.Scan(&titulo)
		default:
			fmt.Println("...")
		}
		titulo = fmt.Sprintf("'%s'", titulo)
		query := fmt.Sprintf("SELECT idproj FROM projetos WHERE %s = %s", variavel, titulo)
		err := db.QueryRow(query).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("SELECT nome, tipo, subtipo, link_circuito FROM projetos WHERE idproj = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&nome, &tipo, &subtipo, &linkCircuito)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Nome: ", nome)
		fmt.Println("Tipo: ", tipo)
		fmt.Println("Subtipo: ", subtipo)
		fmt.Println("Link do circuito: ", linkCircuito)

		fmt.Println("\nVocê deseja:\n1 - Ver peças usadas nesse projeto\n2 - Verificar se você tem todas as peças\n3 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:

		case 2:

		case 3:
			escolha = 4
		}
	}

}

// type Projeto struct {
// 	ID           int    `json:"id"`
// 	Nome         string `json:"nome"`
// 	Tipo         string `json:"tipo"`
// 	LinkCircuito string `json:"link_circuito"`
// 	Descricao    string `json:"descricao"`
