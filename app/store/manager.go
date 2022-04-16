package store

import (
	"database/sql"

	"github.com/google/go-github/v43/github"
)

type Manager struct {
	Database *sql.DB
	GitHub   *github.Client
}
