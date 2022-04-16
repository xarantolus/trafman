package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v43/github"
)

func (m *Manager) StartBackgroundTasks() (err error) {

	go m.runBackgroundTasks()

	return
}

func (m *Manager) runBackgroundTasks() {
	for {
		log.Printf("Running background tasks")

		err := m.runReposUpdate()
		if err != nil {
			log.Printf("[Warning] Running repo data update: %s\n", err.Error())
		}

		nextTime := getNextDay()
		waitDur := time.Until(nextTime)
		if waitDur < 0 {
			panic("invalid nextTime when running background tasks: got " + nextTime.String())
		}

		log.Printf("Running next background tasks on %s", nextTime.String())
		time.Sleep(waitDur)
	}
}

// getNextDay returns a time after midnight
func getNextDay() time.Time {
	now := time.Now().UTC()

	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 5, 0, 0, time.UTC)
}

func getOwnerName(repo *github.Repository) string {
	owner := "unknown"
	if repo.Owner != nil {
		owner = repo.Owner.GetLogin()
	}
	return owner
}

func (m *Manager) runReposUpdate() (err error) {
	ctx := context.Background()

	repos, err := m.fetchAllRepos(ctx)
	if err != nil {
		return
	}

	for _, repo := range repos {
		rerr := m.processRepo(ctx, repo)
		if rerr != nil {
			log.Printf("[Warning] Updating repo data for %s/%s: %s\n", getOwnerName(repo), repo.GetName(), rerr.Error())
		}
	}

	log.Printf("Finished working on %d repos", len(repos))

	return
}

func (m *Manager) fetchAllRepos(ctx context.Context) (repos []*github.Repository, err error) {
	var repoPage = 1

	for {
		log.Printf("Fetching page %d of repositories", repoPage)

		// Get Repos for the currently logged in user
		fetched, resp, err := m.GitHub.Repositories.List(ctx, "", &github.RepositoryListOptions{
			Visibility:  "all",
			Affiliation: "owner,collaborator,organization_member",

			ListOptions: github.ListOptions{
				Page: repoPage,
				// PerPage: 100,
			},
		})
		if err != nil || len(fetched) == 0 {
			return nil, err
		}

		repos = append(repos, fetched...)
		log.Printf("Got %d repos, now have %d", len(fetched), len(repos))

		repoPage = resp.NextPage
		if repoPage == 0 {
			break
		}
	}

	return
}

func (m *Manager) processRepo(ctx context.Context, repo *github.Repository) (err error) {
	var (
		repoUser = getOwnerName(repo)
		repoName = repo.GetName()
	)
	log.Printf("[Background] Working on %s/%s", repoUser, repoName)

	_, err = m.Database.Exec(`insert into Repository(id, username, name, description, is_fork) values ($1, $2, $3, $4, $5)
			on conflict (id) do update set username=EXCLUDED.username, name=EXCLUDED.name, description=EXCLUDED.description, is_fork=EXCLUDED.is_fork`,
		repo.ID, repoUser, repoName, repo.GetDescription(), repo.GetFork())
	if err != nil {
		return fmt.Errorf("inserting basic repo: %s", err.Error())
	}

	_, err = m.Database.Exec(`insert into RepoStats(repo_id, stars, forks, size, subscribers)
								 values ($1, $2, $3, $4, $5)
  							on conflict (repo_id, date) do update set stars=EXCLUDED.stars, forks=EXCLUDED.forks, size=EXCLUDED.size, subscribers=EXCLUDED.subscribers`,
		repo.ID, repo.GetStargazersCount(), repo.GetForksCount(), repo.GetSize(), repo.GetSubscribersCount())
	if err != nil {
		return fmt.Errorf("inserting basic repo info: %s", err.Error())
	}

	views, _, err := m.GitHub.Repositories.ListTrafficViews(ctx, repoUser, repoName, &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		return fmt.Errorf("fetching traffic views: %s", err.Error())
	}
	for _, day := range views.Views {
		_, err = m.Database.Exec(`insert into RepoTrafficViews(repo_id, date, count, uniques)
								 values ($1, $2, $3, $4)
  							on conflict (repo_id, date) do update set count=EXCLUDED.count, uniques=EXCLUDED.uniques`,
			repo.ID, day.GetTimestamp().Time, day.GetCount(), day.GetUniques(),
		)
		if err != nil {
			return fmt.Errorf("inserting traffic view data: %s", err.Error())
		}
	}

	paths, _, err := m.GitHub.Repositories.ListTrafficPaths(ctx, repoUser, repoName)
	if err != nil {
		return fmt.Errorf("fetching traffic paths: %s", err.Error())
	}
	for _, path := range paths {
		_, err = m.Database.Exec(`insert into RepoTrafficPaths(repo_id, path, title, count, uniques)
								 values ($1, $2, $3, $4, $5)
  							on conflict (repo_id, date, path) do update set count=EXCLUDED.count, uniques=EXCLUDED.uniques`,
			repo.ID, path.GetPath(), path.GetTitle(), path.GetCount(), path.GetUniques(),
		)
		if err != nil {
			return fmt.Errorf("inserting traffic path data: %s", err.Error())
		}
	}

	refs, _, err := m.GitHub.Repositories.ListTrafficReferrers(ctx, repoUser, repoName)
	if err != nil {
		return fmt.Errorf("fetching traffic referrers: %s", err.Error())
	}
	for _, ref := range refs {
		_, err = m.Database.Exec(`insert into RepoTrafficReferrers(repo_id, referrer, count, uniques)
								 values ($1, $2, $3, $4)
  							on conflict (repo_id, date, referrer) do update set count=EXCLUDED.count, uniques=EXCLUDED.uniques`,
			repo.ID, ref.GetReferrer(), ref.GetCount(), ref.GetUniques(),
		)
		if err != nil {
			return fmt.Errorf("inserting traffic referrer data: %s", err.Error())
		}
	}

	clones, _, err := m.GitHub.Repositories.ListTrafficClones(ctx, repoUser, repoName, &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		return fmt.Errorf("fetching traffic clones: %s", err.Error())
	}
	for _, clone := range clones.Clones {
		_, err = m.Database.Exec(`insert into RepoTrafficClones(repo_id, date, count, uniques)
								 values ($1, $2, $3, $4)
  							on conflict (repo_id, date) do update set count=EXCLUDED.count, uniques=EXCLUDED.uniques`,
			repo.ID, clone.GetTimestamp().Time, clone.GetCount(), clone.GetUniques(),
		)
		if err != nil {
			return fmt.Errorf("inserting traffic referrer data: %s", err.Error())
		}
	}

	return
}
