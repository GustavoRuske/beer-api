package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GustavoRuske/beer-api/core/beer"
	"github.com/GustavoRuske/beer-api/web/handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	service := beer.NewService(db)

	router := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeBeerHandlers(router, n, service)

	http.Handle("/", router)

	httpService := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err = httpService.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
