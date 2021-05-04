package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoRuske/beer-api/core/beer"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeBeerHandlers(r *mux.Router, n *negroni.Negroni, service beer.UseCase) {
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(getAllBeer(service)),
	)).Methods("GET", "OPTIONS")

	// TO DO criar os outros metodos do UserCase
}

func getAllBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		all, err := service.GetAll()
		if err != nil {
			w.Write(formatJSONError(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if all == nil {
			all = []*beer.Beer{}
		}

		err = json.NewEncoder(w).Encode(all)

		if err != nil {
			w.Write(formatJSONError("Erro ao converter para JSON"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
}
