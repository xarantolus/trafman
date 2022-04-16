package web

import (
	"log"
	"net/http"

	"xarantolus/trafman/config"
	"xarantolus/trafman/store"
)

type Server struct {
	Manager *store.Manager
}

func (s *Server) Run(cfg config.Config) (err error) {

	log.Println("[Server] Start listening on port", cfg.AppPort)
	return http.ListenAndServe(":"+cfg.AppPort, nil)
}
