package server

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
)

// This is the REST handler for the POST /accounts endpoint

type CreateAccountRequest struct {
	Body CreateAccountRequestBody `required:"true"`
}

type CreateAccountRequestBody struct {
	DocumentNumber string `json:"document_number" required:"true" minLength:"1" example:"123456789" doc:"Document number of account"`
}

type CreateAccountResponse struct {
	Body Account
}

func (s *Server) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	// Setup usecase inputs
	input := &accounts.CreateAccountInput{
		DocumentNumber: req.Body.DocumentNumber,
	}
	repo := &accounts.AccountUsecaseRepos{
		Logger:            s.logger,
		AccountsDatastore: s.datastore,
	}

	// Call usecase to create account
	newAccount, err := s.usecases.CreateAccount(ctx, input, repo)

	if err != nil {
		return &CreateAccountResponse{}, DomainToRESTError(err)
	}

	// Return REST representation of response
	return &CreateAccountResponse{
		Body: DomainAccountToREST(newAccount),
	}, nil
}
