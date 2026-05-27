package ui

import (
	"database/sql"
	"fmt"

	"go-index-projetos/api/src/projetos"
)

func MostrarMenuPrinc(db *sql.DB) {
	var escolha int
	for escolha != 3 {
		fmt.Println("Bem-vindo ao sistema de gerenciamento de projetos!\n1 - Cadastrar Projeto\n2 - Buscar Projeto\n3 - Voltar")
		fmt.Scanln(&escolha)
		switch escolha {
		case 1:
			projetos.CadastrarProjeto(db)
		case 2:
			projetos.BuscarProjeto()
		case 3:
			fmt.Println("Voltando ao menu principal...")
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
