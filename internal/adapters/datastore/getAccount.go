package datastore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

func (ds *Datastore) GetAccount (ctx context.Context, accountID string) (*entities.Account, error) {
	var account Account
	err := ds.db.QueryRowxContext(
		ctx,
		"SELECT id,document_num FROM accounts WHERE id = ?",
		accountID,
	).StructScan(&account)
	
	if err != nil {
		// Check if no account was found
		if errors.Is(err, sql.ErrNoRows) {
			ds.logger.Info("No account found in database", "accountID", accountID)
			return nil, entities.NewAccountNotFoundError(accountID)
		} 
		ds.logger.Error("Failed to query database", "error", err.Error())
		return nil, err
	}
	
	domainAccount, err := AccountModelToDomain(account) 
	if err != nil {
		ds.logger.Error("Failed to convert database account to domain model", "error", err.Error())
		return nil, err
	}
	
	return domainAccount, nil
}
