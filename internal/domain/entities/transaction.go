package entities

import (
	"time"

	"github.com/google/uuid"
)

type OperationType int

const (
	UnknownType OperationType = iota
	CashPurchase
	InstalmentPurchase
	Withdrawal
	Payment
)

type Transaction struct {
	ID            uuid.UUID
	AccountID     uuid.UUID
	OperationType OperationType
	Amount        int
	EventDate     time.Time
}
