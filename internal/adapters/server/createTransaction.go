package server

import "context"

type CreateTransactionRequest struct {
}

type CreateTransactionRequestBody struct {
}

type CreateTransactionResponse struct {
}

type CreateTransactionResponseBody struct {
}

func (s *Server) CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return &CreateTransactionResponse{}, UnimplementedErr
}
