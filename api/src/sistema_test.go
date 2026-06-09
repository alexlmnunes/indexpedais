package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFluxoIntegrado(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Erro ao criar o banco de dados simulado: ", err)
	}
	defer db.Close()
}
