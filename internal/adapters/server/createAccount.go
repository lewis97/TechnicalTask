package server

import "context"

type CreateAccountRequest struct {
}

type CreateAccountRequestBody struct {
}

type CreateAccountResponse struct {
}

type CreateAccountResponseBody struct {
}

func (s *Server) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	return &CreateAccountResponse{}, UnimplementedErr
}
