package datastore

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thejerf/slogassert"
)

func Test_DatastoreCreateAccount_Happypath(t *testing.T){
	// Arrange
	mockedDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err, "Failed to create mocked DB")
	sqlxDB := sqlx.NewDb(mockedDB, "mockedDB")
	defer mockedDB.Close()

	// Configure test logger
	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)
	defer logAsserter.AssertEmpty()
	
	ctx := context.Background()
	
	accountID, err := uuid.NewV7()
	require.NoError(t, err)
	
	account := entities.Account{
		ID: accountID,
		DocumentNumber: 123,
		CreatedAt: time.Now(),
	}
	
	expectedSQLstatement := "INSERT INTO accounts (id,document_num,created_at) VALUES ($1,$2,$3)"
	mockSQLResult := sqlmock.NewResult(0, 1)
	mock.
	ExpectExec(expectedSQLstatement).
	WithArgs(account.ID, account.DocumentNumber, account.CreatedAt).
	WillReturnResult(mockSQLResult)
	
	datastore := NewDatastore(sqlxDB, *testLogger)

	// Act
	err = datastore.CreateAccount(ctx, account)
	// Assert
	assert.NoError(t, err)
}

func Test_DatastoreCreateAccount_Errors(t *testing.T){
	// Arrange
	mockedDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err, "Failed to create mocked DB")
	sqlxDB := sqlx.NewDb(mockedDB, "mockedDB")
	defer mockedDB.Close()

	// Configure test logger
	logAsserter := slogassert.New(t, slog.LevelInfo, nil)
	testLogger := slog.New(logAsserter)
	
	ctx := context.Background()
	
	accountID, err := uuid.NewV7()
	require.NoError(t, err)
	
	account := entities.Account{
		ID: accountID,
		DocumentNumber: 123,
		CreatedAt: time.Now(),
	}
	
	expectedSQLstatement := "INSERT INTO accounts (id,document_num,created_at) VALUES ($1,$2,$3)"
	mockSQLErr := errors.New("test error")
	mock.
	ExpectExec(expectedSQLstatement).
	WithArgs(account.ID, account.DocumentNumber, account.CreatedAt).
	WillReturnError(mockSQLErr)
	
	datastore := NewDatastore(sqlxDB, *testLogger)

	// Act
	err = datastore.CreateAccount(ctx, account)
	// Assert
	assert.Equal(t, mockSQLErr, err)
}


