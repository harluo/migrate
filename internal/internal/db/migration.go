package db

import (
	"database/sql"

	"github.com/goexl/db"
	"github.com/harluo/migrate/internal/internal/config"
)

type Migration struct {
	dt     db.Type
	db     *sql.DB
	config *config.Migrate
}

func newMigration(dt db.Type, db *sql.DB, config *config.Migrate) *Migration {
	return &Migration{
		dt:     dt,
		db:     db,
		config: config,
	}
}
