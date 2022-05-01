package store

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/go-github/v43/github"
	"github.com/xarantolus/trafmon/app/web/query"
)

type historicalStats struct {
	Stars int
	Forks int
}

func FetchHistoricalStats(ctx context.Context, client *github.Client, repoUser, repoName string) {
	var evts = map[query.Date]*historicalStats{}

	var setStats = func(date time.Time, s historicalStats) {
		var d = query.NewDate(date)

		se, ok := evts[d]
		if ok {
			if s.Forks != 0 {
				se.Forks += s.Forks
			}
			if s.Stars != 0 {
				se.Stars += s.Stars
			}
		} else {
			se = &s
		}
		evts[d] = se
	}

	stars, err := fetchAllPages(func(opts github.ListOptions) ([]*github.Stargazer, *github.Response, error) {
		return client.Activity.ListStargazers(ctx, repoUser, repoName, &opts)
	}, 100)
	if err != nil {
		return
	}

	for _, s := range stars {
		setStats(s.GetStarredAt().Time.UTC(), historicalStats{Stars: 1})
	}

	forks, err := fetchAllPages(func(opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		return client.Repositories.ListForks(ctx, repoUser, repoName, &github.RepositoryListForksOptions{
			ListOptions: opts,
		})
	}, 100)
	if err != nil {
		return
	}
	for _, f := range forks {
		setStats(f.GetCreatedAt().Time.UTC(), historicalStats{Forks: 1})
	}

	type eventDay struct {
		day  query.Date
		evts historicalStats
	}
	var events []eventDay
	for d, hs := range evts {
		events = append(events, eventDay{day: d, evts: *hs})
	}
	sort.Slice(events, func(i, j int) bool {
		return time.Time(events[i].day).Before(time.Time(events[j].day))
	})

	fmt.Println(len(events))
}

func fetchAllPages[I any](f func(opts github.ListOptions) ([]I, *github.Response, error), maxPerPage int) (result []I, err error) {
	var page = 1

	for {
		fetched, resp, err := f(github.ListOptions{
			Page:    page,
			PerPage: maxPerPage,
		})
		if err != nil || len(fetched) == 0 {
			return result, err
		}

		result = append(result, fetched...)

		page = resp.NextPage
		if page == 0 {
			break
		}
	}

	return
}
