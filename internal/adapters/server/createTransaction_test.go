package server

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/adapters/datastore"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/usecases/transactions"
	mocks "github.com/lewis97/TechnicalTask/mocks/adapters/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thejerf/slogassert"
)

func Test_CreateTransaction(t *testing.T) {
	// Arrange: common
	ctx := context.Background()
	operationID := 1
	amount := 100

	// Configure test logger
	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)

	// Generate uuid for new account
	accountID, err := uuid.NewV7()
	require.NoError(t, err)

	// Generate creation time
	eventDate := time.Now()

	testCases := []struct {
		name             string
		request          CreateTransactionRequest
		setMocks         func(mock *mocks.Usecase)
		expectedResponse *CreateTransactionResponse
		expectedError    error
	}{
		{
			name: "happy path",
			request: CreateTransactionRequest{
				Body: CreateTransactionRequestBody{
					AccountID:     accountID.String(),
					OperationType: operationID,
					Amount:        amount,
				},
			},
			setMocks: func(mock *mocks.Usecase) {
				// Arrange: mock the call to the usecases
				expectedUsecaseInput := &transactions.CreateTransactionInput{
					AccountID:   accountID,
					OperationID: operationID,
					Amount:      amount,
				}
				expectedUsecaseRepo := &transactions.TransactionsUsecaseRepos{
					Logger:                *testLogger,
					AccountsDatastore:     &datastore.Datastore{},
					TransactionsDatastore: &datastore.Datastore{},
				}

				mockTransaction := entities.Transaction{
					ID:            accountID,
					AccountID:     accountID,
					OperationType: entities.OperationType(operationID),
					Amount:        amount,
					EventDate:     eventDate,
				}

				mock.
					EXPECT().
					CreateTransaction(ctx, expectedUsecaseInput, expectedUsecaseRepo).
					Once().
					Return(mockTransaction, nil)
			},
			expectedResponse: &CreateTransactionResponse{
				Body: Transaction{
					ID:            accountID.String(),
					AccountID:     accountID.String(),
					OperationType: operationID,
					Amount:        amount,
					EventDate:     eventDate,
				},
			},
		},
		{
			name: "invalid uuid in request",
			request: CreateTransactionRequest{
				Body: CreateTransactionRequestBody{
					AccountID:     "invalid-id",
					OperationType: operationID,
					Amount:        amount,
				},
			},
			setMocks:         func(mock *mocks.Usecase) {},
			expectedResponse: &CreateTransactionResponse{},
			expectedError:    huma.Error400BadRequest("Invalid account id"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: set usecase mock asserts from testcase
			mockUsecase := mocks.NewUsecase(t)
			tc.setMocks(mockUsecase)

			server := New(Dependencies{
				Logger:    *testLogger,
				Datastore: &datastore.Datastore{},
				Usecases:  mockUsecase,
			})

			// Act
			response, err := server.CreateTransaction(ctx, &tc.request)

			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}

}
