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
		log.Fatalf("Error inserting record: %v", err)
	}

	err = db.QueryRow("SELECT idproj FROM projetos WHERE nome = ?", nome).Scan(&idProj)
	if err != nil {
		log.Fatalf("Error retrieving record: %v", err)
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
				fmt.Println("Essa peça ainda não existe!")
				fmt.Print("Digite o tipo da peça nova: ")
				fmt.Scanln(&tipo)
				fmt.Print("Digite o detalhe da peça nova: ")
				fmt.Scanln(&detalhe)
				fmt.Print("Digite a voltagem da peça nova: ")
				fmt.Scanln(&voltagem)
				fmt.Print("Digite quantidade no estoque da peça nova: ")
				fmt.Scanln(&quant_estoque)
				_, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", tipo, entradaUsuario, detalhe, voltagem, quant_estoque)
				if err != nil {
					log.Fatalf("Error inserting record: %v", err)
				}
				err = db.QueryRow("SELECT LAST_INSERT_ID() FROM pecas").Scan(&idPeca)
				if err != nil {
					log.Fatalf("Error retrieving record: %v", err)
				}
			} else if err != nil {
				log.Fatalf("Error retrieving record: %v", err)
			}
			fmt.Print("Digite a quantidade necessaria para fazer o projeto: ")
			fmt.Scanln(&quant_nec)
			// Conecta o ID do projeto com os IDs de todas as peças que são usadas no projeto
			_, err = db.Exec("INSERT INTO proj_pe (idproj, idpe, quant_pecas) VALUES (?, ?, ?)", idProj, idPeca, quant_nec)
		}
	}
}

func BuscarProjeto() {

}

// type Projeto struct {
// 	ID           int    `json:"id"`
// 	Nome         string `json:"nome"`
// 	Tipo         string `json:"tipo"`
// 	LinkCircuito string `json:"link_circuito"`
// 	Descricao    string `json:"descricao"`
