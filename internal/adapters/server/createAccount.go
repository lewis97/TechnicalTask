package server

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
)

type CreateAccountRequest struct {
	Body CreateAccountRequestBody `required:"true"`
}

type CreateAccountRequestBody struct {
	DocumentNumber int `json:"document_number" required:"true" minLength:"1" minimum:"1" example:"123456789" doc:"Document number of account"`
}

type CreateAccountResponse struct {
	Body Account
}

func (s *Server) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {

	// Setup usecase inputs
	input := &accounts.CreateAccountInput{
		DocumentNumber: uint(req.Body.DocumentNumber),
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

	return &CreateAccountResponse{
		Body: DomainAccountToREST(newAccount),
	}, nil
}
