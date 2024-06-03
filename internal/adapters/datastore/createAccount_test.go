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

func Test_DatastoreCreateAccount(t *testing.T) {
	// Arrange: common variables
	expectedSQLstatement := "INSERT INTO accounts (id,document_num,created_at) VALUES ($1,$2,$3)"
	mockSQLErr := errors.New("test error")

	accountID, err := uuid.NewV7() // generate new uuid for test account
	require.NoError(t, err)

	account := entities.Account{
		ID:             accountID,
		DocumentNumber: 123,
		CreatedAt:      time.Now(),
	}

	testCases := []struct {
		name string
		setDBMock func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "happy path",
			setDBMock: func(mock sqlmock.Sqlmock) {
				// Arrange: mock database operations
				// any successful result is okay (don't rely on this output other than it being a non-error)
				mockSQLResult := sqlmock.NewResult(0, 1) 
				mock.
					ExpectExec(expectedSQLstatement).
					WithArgs(account.ID, account.DocumentNumber, account.CreatedAt).
					WillReturnResult(mockSQLResult)
			},
		},
		{
			name: "database error",
			setDBMock: func(mock sqlmock.Sqlmock) {
				// Arrange: mock database operations
				mock.
					ExpectExec(expectedSQLstatement).
					WithArgs(account.ID, account.DocumentNumber, account.CreatedAt).
					WillReturnError(mockSQLErr)
			},
			expectedError: mockSQLErr,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange: setup mock database
			mockedDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err, "Failed to create mocked DB")
			sqlxDB := sqlx.NewDb(mockedDB, "mockedDB")
			defer mockedDB.Close()

			// mock database calls using test case
			tc.setDBMock(mock)

			// Arrange: configure test logger
			logAsserter := slogassert.New(t, slog.LevelInfo, nil)
			testLogger := slog.New(logAsserter)

			// Arrange setup inputs
			ctx := context.Background()


			datastore := NewDatastore(sqlxDB, *testLogger)

			// Act
			err = datastore.CreateAccount(ctx, account)

			// Assert
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
