package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GustavoRuske/beer-api/core/beer"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeBeerHandlers(r *mux.Router, n *negroni.Negroni, service beer.UseCase) {
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(getAll(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(get(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer", n.With(
		negroni.Wrap(store(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/beer", n.With(
		negroni.Wrap(update(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(remove(service)),
	)).Methods("DELETE", "OPTIONS")
}

func getAll(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		all, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}

		if all == nil {
			all = []*beer.Beer{}
		}

		err = json.NewEncoder(w).Encode(all)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro ao converter para JSON"))
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}

func get(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		b, err := service.Get(id)

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		err = json.NewEncoder(rw).Encode(b)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(formatJSONError("Erro ao converter o JSON"))
			return
		}
	})
}

func store(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var b beer.Beer
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.Store(&b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func update(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		var b beer.Beer
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		if b.ID == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(formatJSONError("invalid ID"))
			return
		}

		err = service.Update(&b)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		rw.WriteHeader(http.StatusOK)
	})
}

func remove(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		_, err = service.Get(id)

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		err = service.Remove(id)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(formatJSONError(err.Error()))
			return
		}

		rw.WriteHeader(http.StatusOK)
	})
}
