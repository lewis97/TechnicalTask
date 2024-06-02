package datastore

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID uuid.UUID `db:"id"`
	DocumentNumber uint `db:"document_num"`
}

type Transaction struct {
	ID uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	OperationID int `db:"operation_id"`
	Amount int `db:"amount"`
	EventTime time.Time `db:"event_time"`
}
