package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request) (err error)

func wrapError(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
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
