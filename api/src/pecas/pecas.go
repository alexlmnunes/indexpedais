package pecas

import (
	"database/sql"
	"fmt"
	"go-index-projetos/api/src/models"
	"log"
)

func CadastrarPeca(db *sql.DB) {

	pe := inputCadastroPeca()
	_, err := db.Exec("INSERT INTO pecas (tipo, valor, detalhe, voltagem, quant_estoque) VALUES (?, ?, ?, ?, ?)", pe.Tipo, pe.Valor, pe.Detalhe, pe.Voltagem, pe.QuantEstoque)
	if err != nil {
		fmt.Println("Erro ao cadastrar peça:", err)
	} else {
		fmt.Println("Peça cadastrada com sucesso!\n")
	}
}

func CadastrarPecaDeProjeto(db *sql.DB, valor string) int {
	pe := inputCadastroPecaComValor(valor)
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
	var pe models.Peca
	termoBusca, colunaBusca := inputBuscarPeca()
	if colunaBusca == "" && termoBusca == "" {
		return
	}

	query := fmt.Sprintf("SELECT idpe, tipo, valor, detalhe, voltagem, quant_estoque FROM pecas WHERE %s = ?", colunaBusca)
	rows, err := db.Query(query, termoBusca)
	if err != nil {
		fmt.Println("Erro ao buscar peça:", err)
	}
	defer rows.Close()

	var listaPecas []models.Peca
	fmt.Print("\n---Peças Encontradas---\n")
	contador := 1
	for rows.Next() {
		err := rows.Scan(&pe.IDpe, &pe.Tipo, &pe.Valor, &pe.Detalhe, &pe.Voltagem, &pe.QuantEstoque)
		if err != nil {
			log.Fatal(err)
		}
		listaPecas = append(listaPecas, pe)
		fmt.Printf("%d - Tipo: %s | Valor: %s | Quantidade em estoque: %d\n", contador, pe.Tipo, pe.Valor, pe.QuantEstoque)
		contador++
	}

	if len(listaPecas) == 0 {
		fmt.Println("Nenhuma peça encontrada...")
		return
	}

	escolhaBusca := inputEscolhaPeca(len(listaPecas))

	pe = listaPecas[escolhaBusca-1]
	idPeca := pe.IDpe
	fmt.Printf("Tipo: %s\nValor: %s\nDetalhe: %s\nVoltagem: %s\nQuantidade em estoque: %d\n",
		pe.Tipo, pe.Valor, pe.Detalhe, pe.Voltagem, pe.QuantEstoque)

	quantidade := inputAlterarEstoqueBusca()

	if quantidade == 0 {
		return
	}
	alterarEstoquePeca(db, idPeca, pe.QuantEstoque+quantidade)
	fmt.Printf("Quantidade atualizada: %d\n", pe.QuantEstoque+quantidade)

}

func alterarEstoquePeca(db *sql.DB, idPeca int, novaQuantidade int) {
	_, err := db.Exec("UPDATE pecas SET quant_estoque = ? WHERE idpe = ?", novaQuantidade, idPeca)
	if err != nil {
		fmt.Println("Erro ao alterar quantidade da peça:", err)
	}
}

func JuntarIdPecaComIdProjeto(db *sql.DB, idProj int, idPeca int) {
	quant_nec := inputQuantidadeNecessaria()

	_, err := db.Exec("INSERT INTO proj_pe (idproj, idpe, quant_pecas) VALUES (?, ?, ?)", idProj, idPeca, quant_nec)
	if err != nil {
		fmt.Println("Erro ao conectar peça com projeto:", err)
	}
}
