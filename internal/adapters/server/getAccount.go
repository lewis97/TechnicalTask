package server

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
)

// This is the REST handler for the GET /accounts{accountID} endpoint

type GetAccountRequest struct {
	AccountID string `path:"accountID" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
}

type GetAccountResponse struct {
	Body Account
}

func (s *Server) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	// convert account ID to uuid ready for the usecase
	accountUUID, err := uuid.Parse(req.AccountID)
	if err != nil {
		s.logger.Error("Unable to parse account UUID in request", "account-id", req.AccountID)
		return &GetAccountResponse{}, huma.Error400BadRequest("Invalid account id")
	}

	// Set up usecase inputs
	input := &accounts.GetAcccountInput{
		AccountID: accountUUID,
	}
	repo := &accounts.AccountUsecaseRepos{
		Logger:            s.logger,
		AccountsDatastore: s.datastore,
	}

	// Get account via usecase
	account, err := s.usecases.GetAccount(ctx, input, repo)
	if err != nil {
		return &GetAccountResponse{}, DomainToRESTError(err)
	}

	// Return REST representation of response
	return &GetAccountResponse{
		Body: DomainAccountToREST(account),
	}, nil
}
