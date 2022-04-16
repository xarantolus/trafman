package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/go-github/v43/github"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"xarantolus/trafman/config"
	"xarantolus/trafman/store"
	"xarantolus/trafman/web"
)

func main() {
	// Set one up @ https://github.com/settings/tokens/new
	cfg, err := config.FromEnvironment()
	if err != nil {
		log.Fatalf("getting config from environment: %s\n", err.Error())
	}

	ctx := context.Background()
	token := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHubToken},
	))
	ghClient := github.NewClient(token)

	// Connect to database
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("connecting to database: %s", err)
	}
	err = database.Ping()
	if err != nil {
		log.Fatalf("pinging database: %s", err.Error())
	}

	var manager = &store.Manager{
		Database: database,
		Github:   ghClient,
	}

	err = manager.StartBackgroundTasks()
	if err != nil {
		log.Fatalf("starting background tasks: %s\n", err.Error())
	}

	defer panic("web server should never stop, but did")
	err = (&web.Server{Manager: manager}).Run(cfg)
	if err != nil {
		log.Fatalf("running web server: %s\n", err.Error())
	}
}
