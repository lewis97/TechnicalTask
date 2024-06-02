package server

import (
	"time"
)

type Account struct {
	ID             string    `json:"account_id" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	DocumentNumber uint      `json:"document_number" example:"123456789" doc:"Document number for account"`
	CreatedAt      time.Time `json:"created_at" example:"2023-12-09T16:09:53.0Z" doc:"Account creation time"`
}

type Transaction struct {
	ID            string    `json:"transaction_id" example:"0124e053-3580-7000-ba62-25ac616ac7f4" doc:"Transaction ID"`
	AccountID     string    `json:"account_id" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	OperationType int       `json:"operation_type" example:"2" doc:"Operation Type"`
	Amount        int       `json:"amount" example:"150" doc:"Transaction amount in the lowest denomination"`
	EventDate     time.Time `json:"event_date" example:"2020-01-05T09:34:18.5893223" doc:"Date and time of transaction"`
}
