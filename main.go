package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v43/github"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"xarantolus/trafman/config"
)

func main() {
	// Set one up @ https://github.com/settings/tokens/new
	cfg, err := config.FromEnvironment()
	if err != nil {
		log.Fatalf("getting config from environment: %s\n", err.Error())
	}
	log.Printf("%#v\n", cfg)

	ctx := context.Background()
	token := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHubToken},
	))

	client := github.NewClient(token)

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("connecting to database: %s", err)
	}

	stats := database.Stats()
	log.Printf("stats: %#v\n", stats)

	rows, err := database.Query("SELECT * from repository")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var repo, user string
		err = rows.Scan(&id, &user, &repo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("id=%d, un=%s, rn=%s", id, repo, user)
	}
	views, _, err := client.Repositories.ListTrafficViews(ctx, "xarantolus", "filtrite", &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", views)

	paths, _, err := client.Repositories.ListTrafficPaths(ctx, "xarantolus", "filtrite")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", paths)

	refs, _, err := client.Repositories.ListTrafficReferrers(ctx, "xarantolus", "filtrite")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", refs)

	clones, _, err := client.Repositories.ListTrafficClones(ctx, "xarantolus", "filtrite", &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", clones)

	log.Println("Start listening on port", cfg.AppPort)
	http.ListenAndServe(":"+cfg.AppPort, nil)
	return
}
