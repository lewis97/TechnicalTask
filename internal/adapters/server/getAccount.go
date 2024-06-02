package server

import "context"

type GetAccountRequest struct {
	AccountID string `path:"accountID" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
}

type GetAccountResponse struct {
	Body Account
}

func (s *Server) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	return &GetAccountResponse{}, UnimplementedErr
}
