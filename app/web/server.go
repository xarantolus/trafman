package web

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xarantolus/trafmon/app/config"
	"github.com/xarantolus/trafmon/app/store"
)

type Server struct {
	Manager *store.Manager

	Frontend embed.FS

	router *mux.Router
}

func (s *Server) Run(cfg config.Config) (err error) {
	s.router = mux.NewRouter()

	api := s.router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/repos", wrapError(s.handleReposAPI))

	// Get the correct frontend path and serve it from every route
	subFS, err := fs.Sub(s.Frontend, "frontend")
	if err != nil {
		return fmt.Errorf("setting up frontend FS: %w", err)
	}

	s.router.PathPrefix("/").HandlerFunc(wrapError(indexHandler(subFS)))

	log.Println("[Server] Start listening on port", cfg.AppPort)
	return http.ListenAndServe(":"+cfg.AppPort, s.router)
}

func indexHandler(subFS fs.FS) handler {
	fileServer := http.FileServer(http.FS(subFS))

	// Basically serve the frontend files, and if we 404 we just serve the vue index
	// it will know what to do because of vue router
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		var resourcePath = strings.TrimPrefix(r.URL.Path, "/")
		_, err = fs.Stat(subFS, resourcePath)
		if err != nil {
			r.URL.Path = "/"
			fileServer.ServeHTTP(w, r)
			return nil
		}
		fileServer.ServeHTTP(w, r)
		return nil
	}
}
