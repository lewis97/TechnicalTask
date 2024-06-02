package server

import "context"

type CreateAccountRequest struct {
	Body CreateAccountRequestBody `required:"true"`
}

type CreateAccountRequestBody struct {
	DocumentNumber int `json:"document_number" required:"true" minLength:"1" example:"123456789" doc:"Document number of account"`
}

type CreateAccountResponse struct {
	Body Account
}

func (s *Server) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	return &CreateAccountResponse{}, UnimplementedErr
}
