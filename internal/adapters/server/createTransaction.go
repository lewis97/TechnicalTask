package server

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/usecases/transactions"
)

type CreateTransactionRequest struct {
	Body CreateTransactionRequestBody
}

type CreateTransactionRequestBody struct {
	AccountID     string `json:"account_id" required:"true" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	OperationType int    `json:"operation_type_id" required:"true" example:"2" doc:"Operation Type"`
	Amount        int    `json:"amount" required:"true" minimum:"1" example:"150" doc:"Transaction amount in the lowest denomination"`
}

type CreateTransactionResponse struct {
	Body Transaction
}

func (s *Server) CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	// convert account ID to uuid
	accountUUID, err := uuid.Parse(req.Body.AccountID)
	if err != nil {
		s.logger.Error("Unable to parse account UUID in request", "account-id", req.Body.AccountID)
		return &CreateTransactionResponse{}, huma.Error400BadRequest("Invalid account id")
	}

	// Setup usecase inputs
	input := &transactions.CreateTransactionInput{
		AccountID:   accountUUID,
		OperationID: req.Body.OperationType,
		Amount:      req.Body.Amount,
	}
	repo := &transactions.TransactionsUsecaseRepos{
		Logger:                s.logger,
		AccountsDatastore:     s.datastore,
		TransactionsDatastore: s.datastore,
	}

	// Call usecase to create transaction
	newTransaction, err := s.usecases.CreateTransaction(ctx, input, repo)

	if err != nil {
		return &CreateTransactionResponse{}, DomainToRESTError(err)
	}

	return &CreateTransactionResponse{
		Body: DomainTransactionToREST(newTransaction),
	}, nil
}
