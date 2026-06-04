package pecas

import (
	"database/sql"
	"fmt"
	"log"
)

func CadastrarPeca(db *sql.DB) {
	var tipo, valor, detalhe, voltagem string
	var quant_estoque int
	fmt.Print("Digite o tipo da peça: ")
	fmt.Scanln(&tipo)
	fmt.Print("Digite o valor da peça: ")
	fmt.Scanln(&valor)
	fmt.Print("Digite o detalhe da peça: ")
	fmt.Scanln(&detalhe)
	fmt.Print("Digite a voltagem da peça: ")
	fmt.Scanln(&voltagem)
	fmt.Print("Digite a quantidade no estoque da peça: ")
	fmt.Scanln(&quant_estoque)
	_, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", tipo, valor, detalhe, voltagem, quant_estoque)
	if err != nil {
		fmt.Println("Erro ao cadastrar peça:", err)
	} else {
		fmt.Println("Peça cadastrada com sucesso!")
	}
}

func BuscarPeca(db *sql.DB) {
	var tipo, valor, detalhe, voltagem, busca, coluna string
	var escolha int
	fmt.Println(("Buscar por:\n1 - tipo\n2 - valor\n3 - detalhe\n4 - voltagem\n5 - voltar"))
	for escolha != 5 {
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			coluna = "tipo"
			fmt.Print("Digite o tipo da peça: ")
			fmt.Scanln(&busca)
		case 2:
			coluna = "valor"
			fmt.Print("Digite o valor da peça: ")
			fmt.Scanln(&busca)
		case 3:
			coluna = "detalhe"
			fmt.Print("Digite o detalhe da peça: ")
			fmt.Scanln(&busca)
		case 4:
			coluna = "voltagem"
			fmt.Print("Digite a voltagem da peça: ")
			fmt.Scanln(&busca)
		default:
			fmt.Println("Opção inválida.")
			return
		}
	}
	busca = fmt.Sprintf("'%s'", busca)
	var idPeca int
	query := fmt.Sprintf("SELECT idpe FROM pecas WHERE %s = %s", coluna, busca)
	err := db.QueryRow(query).Scan(&idPeca)
	if err != nil {
		log.Fatal(err)
	}

}

func alterarEstoquePeca(db *sql.DB, idPeca int, novaQuantidade int) {

}
