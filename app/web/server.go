package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xarantolus/trafmon/app/config"
	"github.com/xarantolus/trafmon/app/store"
)

type Server struct {
	Manager *store.Manager

	router *mux.Router
}

func (s *Server) Run(cfg config.Config) (err error) {
	s.router = mux.NewRouter()

	api := s.router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/repos", wrapError(s.handleReposAPI))

	log.Println("[Server] Start listening on port", cfg.AppPort)
	return http.ListenAndServe(":"+cfg.AppPort, s.router)
}
