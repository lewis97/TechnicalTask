package datastore

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

func (ds *Datastore) GetAccount (ctx context.Context, accountID string) (*entities.Account, error) {
	return &entities.Account{}, nil
}
