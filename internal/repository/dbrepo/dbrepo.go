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

// NewPostgresRepo creates a connection to a PostgreSQL database
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App: a,
		DB: conn,
	}
}