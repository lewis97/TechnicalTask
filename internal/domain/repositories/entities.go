package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

// This contains all the applications interactions with our datastore
type Accounts interface {
	CreateAccount(ctx context.Context, account entities.Account) error
	GetAccount(ctx context.Context, accountID uuid.UUID) (*entities.Account, error)
}

type Transactions interface {
	CreateTransaction(ctx context.Context, transaction entities.Transaction) error
}
