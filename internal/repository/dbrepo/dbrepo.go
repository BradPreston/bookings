package dbrepo

import (
	"database/sql"

	"github.com/bradpreston/bookings/internal/config"
	"github.com/bradpreston/bookings/internal/repository"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

// NewPostgresRepo creates a connection to a PostgreSQL database
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App: a,
		DB: conn,
	}
}

// NewTestingRepo creates a test connection to a PostgreSQL database
func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}