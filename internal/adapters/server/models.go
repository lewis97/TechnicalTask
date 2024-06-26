package server

import (
	"time"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
)

// These are the REST representation of our domain entities
// They are used to marshal/unmarshal data between the client and the app

type Account struct {
	ID             string    `json:"account_id" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	DocumentNumber string    `json:"document_number" example:"123456789" doc:"Document number for account"`
	CreatedAt      time.Time `json:"created_at" example:"2023-12-09T16:09:53.0Z" doc:"Account creation time"`
}

type Transaction struct {
	ID            string    `json:"transaction_id" example:"0124e053-3580-7000-ba62-25ac616ac7f4" doc:"Transaction ID"`
	AccountID     string    `json:"account_id" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	OperationType int       `json:"operation_type" example:"2" doc:"Operation Type"`
	Amount        int       `json:"amount" example:"150" doc:"Transaction amount in the lowest denomination"`
	EventDate     time.Time `json:"event_date" example:"2020-01-05T09:34:18.5893223" doc:"Date and time of transaction"`
}

func DomainAccountToREST(account entities.Account) Account {
	return Account{
		ID:             account.ID.String(),
		DocumentNumber: account.DocumentNumber,
		CreatedAt:      account.CreatedAt,
	}
}

func DomainTransactionToREST(transaction entities.Transaction) Transaction {
	return Transaction{
		ID:            transaction.ID.String(),
		AccountID:     transaction.AccountID.String(),
		OperationType: int(transaction.OperationType),
		Amount:        transaction.Amount,
		EventDate:     transaction.EventDate,
	}
}
