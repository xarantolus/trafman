package store

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/go-github/v43/github"
	"github.com/xarantolus/trafmon/app/web/query"
)

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
