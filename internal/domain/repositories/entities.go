package repositories

import (
	"context"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/google/uuid"
)

// This contains all the applications interactions with our datastore
type Accounts interface {
	CreateAccount(ctx context.Context, account entities.Account) error
	GetAccount(ctx context.Context, accountID uuid.UUID) (*entities.Account, error)
}

type Transactions interface {
	CreateTransaction(ctx context.Context, transaction entities.Transaction) error
}
