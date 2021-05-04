package beer_test

import (
	"database/sql"
	"testing"

	"github.com/GustavoRuske/beer-api/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLagger,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Erro ao conectar no banco de dados: %s", err.Error())
	}

	err = clearDB(db)

	if err != nil {
		t.Fatalf("Erro ao limpar o banco de dados: %s", err.Error())
	}
	defer db.Close()

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Erro ao salvar no banco de dados: %s", err.Error())
	}

	saved, err := service.Get(1)

	if err != nil {
		t.Fatalf("Erro ao buscar no banco de dados: %s", err.Error())
	}

	if saved.ID != 1 {
		t.Fatalf("Dados invalidos. Esperado %d, recebido %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer")

	tx.Commit()
	return err
}

// TO DO - Escrever os testes para os outros metodos
