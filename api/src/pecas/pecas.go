package pecas

import (
	"database/sql"
	"fmt"
	"log"
)

type Peca struct {
	IDpe         int
	Tipo         string
	Valor        string
	Detalhe      string
	Voltagem     string
	QuantEstoque int
}

func CadastrarPeca(db *sql.DB) {
	var pe Peca
	fmt.Print("Digite o tipo da peça: ")
	fmt.Scanln(&pe.Tipo)
	fmt.Print("Digite o valor da peça: ")
	fmt.Scanln(&pe.Valor)
	fmt.Print("Digite o detalhe da peça: ")
	fmt.Scanln(&pe.Detalhe)
	fmt.Print("Digite a voltagem da peça: ")
	fmt.Scanln(&pe.Voltagem)
	fmt.Print("Digite a quantidade no estoque da peça: ")
	fmt.Scanln(&pe.QuantEstoque)
	_, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", pe.Tipo, pe.Valor, pe.Detalhe, pe.Voltagem, pe.QuantEstoque)
	if err != nil {
		fmt.Println("Erro ao cadastrar peça:", err)
	} else {
		fmt.Println("Peça cadastrada com sucesso!")
	}
}

func CadastrarPecaDeProjeto(db *sql.DB, valor string) int {
	var pe Peca
	fmt.Print("Digite o tipo da peça: ")
	fmt.Scanln(&pe.Tipo)
	fmt.Print("Digite o detalhe da peça: ")
	fmt.Scanln(&pe.Detalhe)
	fmt.Print("Digite a voltagem da peça: ")
	fmt.Scanln(&pe.Voltagem)
	fmt.Print("Digite a quantidade no estoque da peça: ")
	fmt.Scanln(&pe.QuantEstoque)
	result, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", pe.Tipo, valor, pe.Detalhe, pe.Voltagem, pe.QuantEstoque)
	if err != nil {
		fmt.Println("Erro ao cadastrar peça:", err)
	}
	fmt.Println("Peça cadastrada com sucesso!")
	idPeca64, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(idPeca64)

}

func BuscarPeca(db *sql.DB) {
	var termoBusca, colunaBusca string
	var escolha, escolhaBusca int
	var pe Peca
	fmt.Println(("Buscar por:\n1 - Tipo\n2 - Valor\n3 - Detalhe\n4 - Voltar"))
	for escolha != 4 {
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			colunaBusca = "tipo"
			fmt.Print("Digite o tipo da peça: ")
			fmt.Scanln(&termoBusca)
		case 2:
			colunaBusca = "valor"
			fmt.Print("Digite o valor da peça: ")
			fmt.Scanln(&termoBusca)
		case 3:
			colunaBusca = "detalhe"
			fmt.Print("Digite o detalhe da peça: ")
			fmt.Scanln(&termoBusca)
		case 4:
			break
		default:
			fmt.Println("Opção inválida.")
			return
		}
		query := fmt.Sprintf("SELECT idpe, tipo, valor, detalhe, voltagem, quant_estoque FROM pecas WHERE %s = ?", colunaBusca)
		rows, err := db.Query(query, termoBusca)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var listaPecas []Peca
		fmt.Print("\n---Peças Encontradas---\n")
		contador := 1
		for rows.Next() {
			err := rows.Scan(&pe.IDpe, &pe.Tipo, &pe.Valor, &pe.Detalhe, &pe.Voltagem, &pe.QuantEstoque)
			if err != nil {
				log.Fatal(err)
			}
			listaPecas = append(listaPecas, pe)
			fmt.Printf("%d - Tipo: %s | Valor: %s\n", contador, pe.Tipo, pe.Valor)
			contador++
		}

		if len(listaPecas) == 0 {
			fmt.Println("Nenhuma peça encontrada...")
			break
		}
		fmt.Println("Digite o número da peça que você deseja: ")
		fmt.Scan(&escolhaBusca)
		if escolhaBusca < 1 || escolhaBusca > len(listaPecas) {
			fmt.Println("Número inválido...")
			break
		}

		pe := listaPecas[escolhaBusca-1]
		idPeca := pe.IDpe
		fmt.Printf("Tipo: %s\nValor: %s\nDetalhe: %s\nVoltagem: %s\nQuantidade em estoque: %d\n",
			pe.Tipo, pe.Valor, pe.Detalhe, pe.Voltagem, pe.QuantEstoque)

		for escolha != 4 {
			fmt.Println("\nVocê deseja:\n1 - Aumentar quantidade em estoque\n2 - Reduzir quantidade em estoque\n3 - Voltar")
			fmt.Scan(&escolha)
			switch escolha {
			case 1:
				var somaQuantidade int
				fmt.Print("Digite a quantidade a ser adicionada: ")
				fmt.Scan(&somaQuantidade)
				pe.QuantEstoque += somaQuantidade
			case 2:
				var subtraiQuantidade int
				fmt.Print("Digite a quantidade a ser removida: ")
				fmt.Scan(&subtraiQuantidade)
				pe.QuantEstoque -= subtraiQuantidade
			case 3:
				break
			default:
				fmt.Println("Opção inválida.")
				return
			}
			alterarEstoquePeca(db, idPeca, pe.QuantEstoque)
			fmt.Printf("Quantidade atualizada: %d\n", pe.QuantEstoque)
		}

	}

}

func alterarEstoquePeca(db *sql.DB, idPeca int, novaQuantidade int) {
	_, err := db.Exec("UPDATE pecas SET quant_estoque = ? WHERE idpe = ?", novaQuantidade, idPeca)
	if err != nil {
		fmt.Println("Erro ao alterar quantidade da peça:", err)
	}
}
