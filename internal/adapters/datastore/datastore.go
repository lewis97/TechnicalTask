package datastore

import "github.com/jmoiron/sqlx"

type Datastore struct {
	db *sqlx.DB
}

func NewDatastore(db *sqlx.DB) *Datastore {
	return &Datastore{
		db: db,
	}
}
