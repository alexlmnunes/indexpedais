package projetos

import (
	"fmt"
	"go-index-projetos/api/src/models"
)

func inputProjeto() models.Projeto {
	var p models.Projeto
	fmt.Print("Digite o nome do projeto: ")
	fmt.Scan(&p.Nome)

	fmt.Print("Digite o tipo do projeto: ")
	fmt.Scan(&p.Tipo)

	fmt.Print("Digite o subtipo do projeto: ")
	fmt.Scan(&p.Subtipo)

	fmt.Print("Digite o link do circuito: ")
	fmt.Scan(&p.LinkCircuito)
	return p
}

func inputBuscarValorPeca(i int) string {
	var entradaUsuario string
	fmt.Print("Digite o valor da ", i, "a peça (ou 'sair' caso não tenha mais peças):")
	fmt.Scan(&entradaUsuario)
	return entradaUsuario
}

func inputBuscarProjeto() (string, string) {
	var termoBusca, colunaBusca string
	var escolha int
	for {
		fmt.Println("Buscar por: \n1 - Nome\n2 - Tipo\n3 - Link\n4 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			colunaBusca = "nome"
			fmt.Println("Digite o nome do projeto: ")
			fmt.Scan(&termoBusca)
			return colunaBusca, termoBusca
		case 2:
			colunaBusca = "tipo"
			fmt.Println("Digite o tipo de projeto: ")
			fmt.Scan(&termoBusca)
			return colunaBusca, termoBusca
		case 3:
			colunaBusca = "link"
			fmt.Println("Digite o link do projeto: ")
			fmt.Scan(&termoBusca)
			return colunaBusca, termoBusca
		case 4:
			return "", ""
		default:
			fmt.Println("Escolha inválida...")
		}

	}
}

func inputEscolhaProjeto(tamanhoLista int) int {
	var escolhaBusca int

	for {
		fmt.Println("Digite o número do projeto que você deseja: ")
		fmt.Scan(&escolhaBusca)
		if escolhaBusca >= 1 && escolhaBusca <= tamanhoLista {
			return escolhaBusca
		}
		fmt.Println("Número inválido...")
	}
}

func inputVerificarPecasOuEstoque() int {
	var escolha int

	for {
		fmt.Println("\nVocê deseja:\n1 - Ver peças usadas nesse projeto\n2 - Verificar se você tem todas as peças\n3 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 0
		default:
			fmt.Println("Opção inválida.")
		}

	}

}

func inputBaixaEstoque() bool {
	var escolhaBaixa int
	for {
		fmt.Println("Você tem todas as peças necessárias para esse projeto! Deseja dar baixa no estoque e fazer o projeto?\n1 - Sim\n2 - Não")
		fmt.Scan(&escolhaBaixa)
		switch escolhaBaixa {
		case 1:
			return true
		case 2:
			return false
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
