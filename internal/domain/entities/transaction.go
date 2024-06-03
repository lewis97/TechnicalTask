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

// Valid operation type is valid (1,2,3,4 only)
func ValidateOperationType(ot int) bool {
	return ot > 0 && ot <= 4
}

// Convert operation type enum to its string representation 
func (ot OperationType) String() string {
	switch ot {
	case UnknownType:
		return "INVALID"
	case CashPurchase:
		return "CASH PURCHASE"
	case InstalmentPurchase:
		return "INSTALLMENT PURCHASE"
	case Withdrawal:
		return "WITHDRAWAL"
	case Payment:
		return "PAYMENT"
	default:
		return "INVALID"
	}
}
