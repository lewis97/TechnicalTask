package entities

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID             uuid.UUID
	DocumentNumber string
	CreatedAt      time.Time
}
