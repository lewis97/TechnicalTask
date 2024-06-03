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

func Test_CreateAccount_HappyPath(t *testing.T) {
	// Arrange
	ctx := context.Background()
	docNum := "123"

	// Arrange: generate new uuid for account
	accountID, err := uuid.NewV7()
	require.NoError(t, err)

	// Arrange: mock uuid generation in usecase
	uuidGenMock := uuidMock.NewUUIDGenerator(t)
	uuidGenMock.EXPECT().New().Once().Return(accountID, nil)

	// Arrange: mock clock for time.Now() calls
	createdAt := time.Now()
	clockMock := clockMock.NewClock(t)
	clockMock.EXPECT().Now().Once().Return(createdAt)

	accountsUsecase := NewAccountsUsecase(uuidGenMock, clockMock)

	// Arrange: setup logger
	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)

	// Arrange: mock datastore operations
	datastoreMock := repoMock.NewAccounts(t)

	// mock returning account not found error from database so the account can be created
	datastoreMock.
		EXPECT().
		GetAccountByDoc(ctx, docNum).
		Once().
		Return(&entities.Account{}, entities.NewAccountNotFoundByDocNumError(docNum))

	// mock create account call to database - return nil (no error)
	expectedAccount := entities.Account{
		ID:             accountID,
		DocumentNumber: docNum,
		CreatedAt:      createdAt,
	}
	datastoreMock.EXPECT().CreateAccount(ctx, expectedAccount).Once().Return(nil)

	accountRepos := AccountUsecaseRepos{
		Logger:            *testLogger,
		AccountsDatastore: datastoreMock,
	}

	createAccountInput := CreateAccountInput{
		DocumentNumber: docNum,
	}

	// Act
	response, err := accountsUsecase.CreateAccount(ctx, &createAccountInput, &accountRepos)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, response)
}
