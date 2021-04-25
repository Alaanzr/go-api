package tweet

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type HttpHandler func(string, func(w http.ResponseWriter, r *http.Request))

var tweets = []Tweet{
	{Content: "tweet 1"},
	{Content: "tweet 2"},
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tweets)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tweets[0])
}

func create(w http.ResponseWriter, r *http.Request) {
	tweet := &Tweet{}
	
	json.NewDecoder(r.Body).Decode(tweet)

	v := validator.New()

	if err := v.Struct(tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweets = append(tweets, *tweet)

	w.WriteHeader(http.StatusCreated)
}

func RegisterHandlers(router *chi.Mux) {
	router.Get("/api/tweets", getAll)
	router.Get("/api/tweets/{id}", getOne)
	router.Post("/api/tweets", create)
}
