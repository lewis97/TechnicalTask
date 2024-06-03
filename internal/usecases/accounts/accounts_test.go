package accounts

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	repoMock "github.com/lewis97/TechnicalTask/mocks/domain/repositories"
	clockMock "github.com/lewis97/TechnicalTask/mocks/drivers/clock"
	uuidMock "github.com/lewis97/TechnicalTask/mocks/drivers/uuidgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thejerf/slogassert"
)

func Test_CreateAccount_HappyPath(t *testing.T){
	// Arrange
	ctx := context.Background()
	docNum := 123

	// Generate new uuid for account
	accountID, err := uuid.NewV7()
	require.NoError(t, err)

	// mock uuid generation in usecase
	uuidGenMock := uuidMock.NewUUIDGenerator(t)
	uuidGenMock.EXPECT().New().Once().Return(accountID, nil)

	// mock clock for time.Now() calls
	createdAt := time.Now()
	clockMock := clockMock.NewClock(t)
	clockMock.EXPECT().Now().Once().Return(createdAt)

	accountsUsecase := NewAccountsUsecase(uuidGenMock, clockMock)

	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)

	// Mock datastore create account
	datastoreMock := repoMock.NewAccounts(t)
	expectedAccount := entities.Account{
		ID:             accountID,
		DocumentNumber: uint(docNum),
		CreatedAt:      createdAt,
	}
	datastoreMock.EXPECT().CreateAccount(ctx, expectedAccount).Once().Return(nil)

	accountRepos := AccountUsecaseRepos{
		Logger: *testLogger,
		AccountsDatastore: datastoreMock,
	}

	createAccountInput := CreateAccountInput{
		DocumentNumber: uint(docNum),
	}

	// Act
	response, err := accountsUsecase.CreateAccount(ctx, &createAccountInput, &accountRepos)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, response)
}
