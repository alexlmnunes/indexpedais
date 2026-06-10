package ui

import (
	"database/sql"
	"fmt"
	"go-index-projetos/api/src/pecas"
	"go-index-projetos/api/src/projetos"
)

func MostrarMenuPrinc(db *sql.DB) {
	var escolha int
	for escolha != 3 {
		fmt.Println("Bem-vindo ao sistema de gerenciamento de projetos e peças!\n1 - Gerenciar Projetos\n2 - Gerenciar Peças\n3 - Sair")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			MostrarMenuProjetos(db)
		case 2:
			MostrarMenuPecas(db)
		case 3:
			fmt.Println("Saindo do sistema...")
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func MostrarMenuProjetos(db *sql.DB) {
	var escolha int
	for escolha != 3 {
		fmt.Println("Gerenciamento de projetos!\n1 - Cadastrar Projeto\n2 - Buscar Projeto\n3 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			projetos.CadastrarProjeto(db)
		case 2:
			projetos.BuscarProjeto(db)
		case 3:
			fmt.Println("Voltando ao menu principal...")
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func MostrarMenuPecas(db *sql.DB) {
	var escolha int
	for escolha != 3 {
		fmt.Println("Gerenciamento de peças!\n1 - Cadastrar Peça\n2 - Buscar Peça\n3 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			pecas.CadastrarPeca(db)
		case 2:
			pecas.BuscarPeca(db)
		case 3:
			fmt.Println("Voltando ao menu principal...")
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
