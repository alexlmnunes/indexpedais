package projetos

import (
	"database/sql"
	"errors"
	"fmt"
	"go-index-projetos/api/src/models"
	"go-index-projetos/api/src/pecas"
	"log"
)

func CadastrarProjeto(db *sql.DB) {
	// Cadastrar dados que vao para a tabela projetos
	p := inputProjeto()
	var idPeca int

	result, err := db.Exec("INSERT INTO projetos (nome, tipo, subtipo, link_circuito) VALUES (?, ?, ?, ?)", p.Nome, p.Tipo, p.Subtipo, p.LinkCircuito)
	if err != nil {
		fmt.Println("Erro ao cadastrar projeto: ", err)
	}

	id64, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao obter ID do projeto cadastrado: ", err)
	}

	p.ID = int(id64)

	// Loop para receber as peças do projeto, parando para cadastrar uma peça quando o mesmo ainda nao existir na tabela pecas, e conectando o id do projeto com o id do peça na tabela proj_pe
	for i := 1; ; i++ {
		valorPeca := inputBuscarValorPeca(i)
		if valorPeca != "sair" {
			err := db.QueryRow("SELECT idpe FROM pecas WHERE valor = ?", valorPeca).Scan(&idPeca)
			// Verifica se essa peça já existe na tabela ou se precisa cadastrar
			if errors.Is(err, sql.ErrNoRows) {
				idPeca = pecas.CadastrarPecaDeProjeto(db, valorPeca)
			} else if err != nil {
				fmt.Println("Erro ao buscar peça: ", err)
			}
			pecas.JuntarIdPecaComIdProjeto(db, p.ID, idPeca)
		}
	}
}

func BuscarProjeto(db *sql.DB) {
	var p models.Projeto

	colunaBusca, termoBusca := inputBuscarProjeto()

	if colunaBusca == "" && termoBusca == "" {
		return
	}

	query := fmt.Sprintf("SELECT idproj, nome, tipo, subtipo, link_circuito FROM projetos WHERE %s = ?", colunaBusca)
	rows, err := db.Query(query, termoBusca)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var listaProjetos []models.Projeto
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
		return
	}

	escolhaBusca := inputEscolhaProjeto(len(listaProjetos))
	projetoSelecionado := listaProjetos[escolhaBusca-1]
	idproj := projetoSelecionado.ID

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nome: %s\nTipo: %s\nSubtipo: %s\nLink do circuito: %s\n",
		projetoSelecionado.Nome, projetoSelecionado.Tipo, projetoSelecionado.Subtipo, projetoSelecionado.LinkCircuito)

	escolha := inputVerificarPecasOuEstoque()
	if escolha == 0 {
		return
	}
	switch escolha {
	case 1:
		verificarPecasProjeto(db, projetoSelecionado)
	case 2:
		if compararQuantEstoque(db, idproj) {
			if inputBaixaEstoque() {
				baixaQuantEstoque(db, idproj)
				fmt.Print("Pronto!")
			}
		} else {
			fmt.Println("Você não tem todas as peças necessárias para esse projeto.")
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

func verificarPecasProjeto(db *sql.DB, p models.Projeto) {
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
