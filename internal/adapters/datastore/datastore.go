package datastore

import (
	"log/slog"
	"github.com/jmoiron/sqlx"
)

type Datastore struct {
	db *sqlx.DB
	logger slog.Logger
}

func NewDatastore(db *sqlx.DB, logger slog.Logger) *Datastore {
	return &Datastore{
		db: db,
		logger: logger,
	}
}
