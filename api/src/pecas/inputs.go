package pecas

import (
	"fmt"
	"go-index-projetos/src/models"
)

func inputCadastroPeca() models.Peca {
	var pe models.Peca
	fmt.Print("Digite o tipo da peça: ")
	fmt.Scan(&pe.Tipo)
	fmt.Print("Digite o valor da peça: ")
	fmt.Scan(&pe.Valor)
	fmt.Print("Digite o detalhe da peça: ")
	fmt.Scan(&pe.Detalhe)
	fmt.Print("Digite a voltagem da peça: ")
	fmt.Scan(&pe.Voltagem)
	fmt.Print("Digite a quantidade no estoque da peça: ")
	fmt.Scan(&pe.QuantEstoque)
	return pe
}

func inputCadastroPecaComValor(valor string) models.Peca {
	var pe models.Peca
	pe.Valor = valor
	fmt.Print("Digite o tipo da peça: ")
	fmt.Scan(&pe.Tipo)
	fmt.Print("Digite o detalhe da peça: ")
	fmt.Scan(&pe.Detalhe)
	fmt.Print("Digite a voltagem da peça: ")
	fmt.Scan(&pe.Voltagem)
	fmt.Print("Digite a quantidade no estoque da peça: ")
	fmt.Scan(&pe.QuantEstoque)
	return pe
}

func inputBuscarPeca() (string, string) {
	var escolha int
	var termoBusca, colunaBusca string

	for {
		fmt.Println("Buscar por:\n1 - Tipo\n2 - Valor\n3 - Detalhe\n4 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			colunaBusca = "tipo"
			fmt.Print("Digite o tipo da peça: ")
			fmt.Scan(&termoBusca)
			return termoBusca, colunaBusca
		case 2:
			colunaBusca = "valor"
			fmt.Print("Digite o valor da peça: ")
			fmt.Scan(&termoBusca)
			return termoBusca, colunaBusca
		case 3:
			colunaBusca = "detalhe"
			fmt.Print("Digite o detalhe da peça: ")
			fmt.Scan(&termoBusca)
			return termoBusca, colunaBusca
		case 4:
			return "", ""
		default:
			fmt.Println("Opção inválida.")
		}

	}
}

func inputEscolhaPeca(tamanhoLista int) int {
	var escolhaBusca int
	for {
		fmt.Println("Digite o número da peça que você deseja: ")
		fmt.Scan(&escolhaBusca)
		if escolhaBusca > 0 && escolhaBusca <= tamanhoLista {
			return escolhaBusca

		}
		fmt.Println("Número inválido...\nPor favor, digite um número entre 1 e", tamanhoLista)
	}
}

func inputAlterarEstoqueBusca() int {
	var escolha, quantidade int
	for {
		fmt.Println("\nVocê deseja:\n1 - Aumentar quantidade em estoque\n2 - Reduzir quantidade em estoque\n3 - Voltar")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			fmt.Print("Digite a quantidade a ser adicionada: ")
			fmt.Scan(&quantidade)
			return quantidade
		case 2:
			fmt.Print("Digite a quantidade a ser removida: ")
			fmt.Scan(&quantidade)
			return -quantidade
		case 3:
			return 0
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func inputQuantidadeNecessaria() int {
	var quant_nec int
	fmt.Print("Digite a quantidade necessaria para fazer o projeto: ")
	fmt.Scan(&quant_nec)
	return quant_nec
}
