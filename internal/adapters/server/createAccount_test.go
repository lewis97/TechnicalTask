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
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
	mocks "github.com/lewis97/TechnicalTask/mocks/adapters/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thejerf/slogassert"
)

func Test_CreateAccount(t *testing.T) {
	// Arrange: common
	ctx := context.Background()
	docNumber := "123"

	// Configure test logger
	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)
	defer logAsserter.AssertEmpty()

	// Generate uuid for new account
	uuid, err := uuid.NewV7()
	require.NoError(t, err)

	// Generate creation time
	createdAt := time.Now()

	testCases := []struct {
		name             string
		mockAccount      entities.Account
		mockError        error
		expectedResponse *CreateAccountResponse
		expectedError    error
	}{
		{
			name: "happy path",
			mockAccount: entities.Account{
				ID:             uuid,
				DocumentNumber: docNumber,
				CreatedAt:      createdAt,
			},
			mockError: nil,
			expectedResponse: &CreateAccountResponse{
				Body: Account{
					ID:             uuid.String(),
					DocumentNumber: docNumber,
					CreatedAt:      createdAt,
				},
			},
			expectedError: nil,
		},
		{
			name:             "create account usecase error",
			mockAccount:      entities.Account{},
			mockError:        entities.NewInvalidInputError("test error"),
			expectedResponse: &CreateAccountResponse{},
			expectedError:    huma.Error400BadRequest("test error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockUsecase := mocks.NewUsecase(t)

			// Create server with injected mocks
			server := New(Dependencies{
				Logger:    *testLogger,
				Datastore: &datastore.Datastore{},
				Usecases:  mockUsecase,
			})

			// Arrange: setup request
			request := CreateAccountRequest{
				Body: CreateAccountRequestBody{
					DocumentNumber: docNumber,
				},
			}

			// Arrange: mock the call to the usecases
			expectedUsecaseInput := &accounts.CreateAccountInput{
				DocumentNumber: docNumber,
			}
			expectedUsecaseRepo := &accounts.AccountUsecaseRepos{
				Logger:            *testLogger,
				AccountsDatastore: &datastore.Datastore{},
			}

			mockUsecase.
				EXPECT().
				CreateAccount(ctx, expectedUsecaseInput, expectedUsecaseRepo).
				Once().
				Return(tc.mockAccount, tc.mockError)

			// Act
			response, err := server.CreateAccount(ctx, &request)

			// Assert
			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}

}
