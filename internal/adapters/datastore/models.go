package datastore

import (
	"time"

	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

type Account struct {
	ID             string    `db:"id"`
	DocumentNumber uint      `db:"document_num"`
	CreatedAt      time.Time `db:"created_at"`
}

type Transaction struct {
	ID          string    `db:"id"`
	AccountID   string    `db:"account_id"`
	OperationID int       `db:"operation_id"`
	Amount      int       `db:"amount"`
	EventTime   time.Time `db:"event_time"`
}

func AccountModelToDomain(account Account) (*entities.Account, error) {
	id, err := uuid.Parse(account.ID)
	if err != nil {
		return nil, err
	}
	return &entities.Account{
		ID:             id,
		DocumentNumber: account.DocumentNumber,
		CreatedAt:      account.CreatedAt,
	}, nil
}
