package server

import "context"

type CreateTransactionRequest struct {
	Body CreateAccountRequestBody
}

type CreateTransactionRequestBody struct {
	AccountID     string `json:"account_id" required:"true" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f" doc:"Account ID"`
	OperationType int    `json:"operation_type_id" required:"true" example:"2" doc:"Operation Type"`
	Amount        int    `json:"amount" required:"true" example:"150" doc:"Transaction amount in the lowest denomination"`
}

type CreateTransactionResponse struct {
	Body Transaction
}

func (s *Server) CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return &CreateTransactionResponse{}, UnimplementedErr
}
