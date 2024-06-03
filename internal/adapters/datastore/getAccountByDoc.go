package datastore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

// Queries the database for a given account by the accounts ID
func (ds *Datastore) GetAccountByDoc(ctx context.Context, documentNumber uint) (*entities.Account, error) {
	var account Account // will hold the account record if found

	// Query database and store the fetched row (account) in account struct
	err := ds.db.QueryRowxContext(
		ctx,
		"SELECT id,document_num,created_at FROM accounts WHERE document_num = $1",
		documentNumber,
	).StructScan(&account)

	if err != nil {
		// Check if no account was found
		if errors.Is(err, sql.ErrNoRows) {
			ds.logger.Info("No account found in database", "documentNumber", documentNumber)
			return nil, entities.NewAccountNotFoundByDocNumError(documentNumber)
		}
		// Error is unexpected & needs to be handled accordingly
		ds.logger.Error("Failed to query database", "error", err.Error())
		return nil, err
	}

	// Convert the found account into our domain object
	domainAccount, err := AccountModelToDomain(account)
	if err != nil {
		ds.logger.Error("Failed to convert database account to domain model", "error", err.Error())
		return nil, err
	}

	return domainAccount, nil
}
