package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func main() {
	// Set one up @ https://github.com/settings/tokens/new
	var (
		ghToken = os.Getenv("GITHUB_TOKEN")
	)
	if strings.TrimSpace(ghToken) == "" {
		log.Fatalf("no GITHUB_TOKEN env variable available")
	}

	ctx := context.Background()
	token := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	))

	client := github.NewClient(token)

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

	http.ListenAndServe(":2000", nil)
	return
}
