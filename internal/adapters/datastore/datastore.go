package datastore

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Datastore struct {
	db     *sqlx.DB
	logger slog.Logger
}

func NewDatastore(db *sqlx.DB, logger slog.Logger) *Datastore {
	return &Datastore{
		db:     db,
		logger: logger,
	}
}
