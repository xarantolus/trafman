package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	errNotFound = fmt.Errorf("404 Not Found")
)

type handler func(w http.ResponseWriter, r *http.Request) (err error)

func wrapError(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		if errors.Is(err, errNotFound) {
			http.Error(w, errNotFound.Error(), http.StatusNotFound)
			return
		}

		log.Printf("[Error] serving %q: %s\n", r.URL.Path, err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func serveJSON(w http.ResponseWriter, r *http.Request, data any) (err error) {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
