package tweet

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type HttpHandler func(string, func(w http.ResponseWriter, r *http.Request))

func getAll(w http.ResponseWriter, r *http.Request) {
	tweets, err := getTweets()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(tweets)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	tweets, err := getTweets()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(tweets)
}

func create(w http.ResponseWriter, r *http.Request) {
	tweet := &Tweet{}
	
	json.NewDecoder(r.Body).Decode(tweet)

	v := validator.New()

	if err := v.Struct(tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := createTweet(*tweet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func RegisterHandlers(router *chi.Mux) {
	router.Get("/api/tweets", getAll)
	router.Get("/api/tweets/{id}", getOne)
	router.Post("/api/tweets", create)
}
