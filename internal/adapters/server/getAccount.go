package server

import "context"

type GetAccountRequest struct {
}

type GetAccountRequestBody struct {
}

type GetAccountResponse struct {
}

type GetAccountResponseBody struct {
}

func (s *Server) GetAccount(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	return &GetAccountResponse{}, UnimplementedErr
}
