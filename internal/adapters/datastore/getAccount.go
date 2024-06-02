package datastore

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

func (ds *Datastore) GetAccount (ctx context.Context, accountID uuid.UUID) (*entities.Account, error) {
	// Convert acountID to string
	id := accountID.String()
	
	var account Account
	err := ds.db.QueryRowxContext(
		ctx,
		"SELECT id,document_num FROM accounts WHERE id = $1",
		id,
	).StructScan(&account)
	
	if err != nil {
		// Check if no account was found
		if errors.Is(err, sql.ErrNoRows) {
			ds.logger.Info("No account found in database", "accountID", accountID)
			return nil, entities.NewAccountNotFoundError(id)
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
