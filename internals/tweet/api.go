package tweet

import (
	"encoding/json"
	"net/http"
)

type HttpHandler func(string, func(w http.ResponseWriter, r *http.Request))

var tweets = []Tweet{
	Tweet{Content: "tweet 1"},
	Tweet{Content: "tweet 2"},
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tweets)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tweets[0])
}

func RegisterHandlers(handler HttpHandler) {
	handler("/api/tweets", getAll)
	handler("/api/tweets/{id}", getOne)
}
