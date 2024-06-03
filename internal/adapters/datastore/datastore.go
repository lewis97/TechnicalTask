package datastore

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

// Datastore is used to interact with the database and implements our domain repository operations
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
