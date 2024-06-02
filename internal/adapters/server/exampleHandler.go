package server

import (
	"context"
	"fmt"
)

type TestRequest struct {
}

type TestRequestBody struct {
}

type TestResponse struct {
	Body TestResponseBody
}

type TestResponseBody struct {
	Msg string `json:"msg" example:"helo doc:"msg"`
}

func (s *Server) TestHandler(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	// name := req.Body.Name
	respBody := TestResponseBody{
		Msg: fmt.Sprintf("Hello"), // %s", name),
	}
	return &TestResponse{
		Body: respBody,
	}, nil
}
