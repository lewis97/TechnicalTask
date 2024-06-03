package datastore

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

// Creates an account in the database, returns nil if successful
func (ds *Datastore) CreateAccount(ctx context.Context, account entities.Account) error {
	_, err := ds.db.ExecContext(
		ctx,
		"INSERT INTO accounts (id,document_num,created_at) VALUES ($1,$2,$3)",
		account.ID,
		account.DocumentNumber,
		account.CreatedAt,
	)

	if err != nil {
		ds.logger.Error("failed to create account in db", "error", err.Error())
	}

	return err

}
