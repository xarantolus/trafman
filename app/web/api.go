package web

import (
	"net/http"
	"time"
)

func (s *Server) handleReposAPI(w http.ResponseWriter, r *http.Request) (err error) {
	type repository struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		Name        string `json:"name"`
		Description string `json:"description"`
		IsFork      bool   `json:"is_fork"`

		DownloadCount int `json:"download_count"`

		LastUpdate time.Time `json:"last_update"`

		Stars int `json:"stars"`
		Forks int `json:"forks"`
		Size  int `json:"size"`
	}

	rows, err := s.Manager.Database.Query(`
	WITH repos AS (
		SELECT DISTINCT ON (r.id) r.id, r.username, r.name, r.description, r.is_fork, s.stars, s.forks, s.size, s.date_time, coalesce(sum(ra.download_count), 0) as download_count
			FROM Repositories r
			join RepoStats s on r.id = s.repo_id
			left join releases rs on r.id = rs.repo_id
			left join releaseassets ra on rs.id = ra.release_id
		group by r.id, r.username, r.name, r.description, r.is_fork, s.stars, s.forks, s.size, s.date_time
		ORDER BY r.id, s.date_time DESC
	)
	SELECT * FROM repos r
	ORDER BY (r.download_count+r.stars+r.forks) DESC`)
	if err != nil {
		return
	}
	defer rows.Close()

	var result []repository
	for rows.Next() {
		var repo = repository{}

		err = rows.Scan(&repo.ID, &repo.Username, &repo.Name, &repo.Description, &repo.IsFork, &repo.Stars, &repo.Forks, &repo.Size, &repo.LastUpdate, &repo.DownloadCount)
		if err != nil {
			return
		}

		result = append(result, repo)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return serveJSON(w, r, result)
}
