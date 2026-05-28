package main

import (
	"database/sql"
	"fmt"
	"go-index-projetos/api/src/ui"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Lixo123@tcp(127.0.0.1:3306)/indexprojetos")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Banco de dados inicializado\n.\n.\n.")
	ui.MostrarMenuPrinc(db)
}
