package transactions

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/domain/repositories"
	"github.com/lewis97/TechnicalTask/internal/drivers/clock"
	"github.com/lewis97/TechnicalTask/internal/drivers/uuidgen"
)

type TransactionsUsecase struct {
	uuidGen uuidgen.UUIDGenerator
	clock clock.Clock
}

type TransactionsUsecaseRepos struct {
	Logger                slog.Logger
	AccountsDatastore     repositories.Accounts
	TransactionsDatastore repositories.Transactions
}

func NewTransactionUsecase(uuidGenerator uuidgen.UUIDGenerator, clock clock.Clock) *TransactionsUsecase {
	return &TransactionsUsecase{
		uuidGen: uuidGenerator,
		clock: clock,
	}
}

// Create a transaction

type CreateTransactionInput struct {
	AccountID   uuid.UUID
	OperationID int
	Amount      int
}

func (input *CreateTransactionInput) Validate() error {
	if !entities.ValidateOperationType(input.OperationID) {
		return entities.NewInvalidInputError("Invalid operation type id")
	}
	if input.Amount <= 0 {
		return entities.NewInvalidInputError("Invalid amount - must be >0")
	}
	return nil
}

func (tc *TransactionsUsecase) CreateTransaction(ctx context.Context, input *CreateTransactionInput, repo *TransactionsUsecaseRepos) (entities.Transaction, error) {
	// Validate the input
	err := input.Validate()
	if err != nil {
		repo.Logger.Error("Create transaction input failed validation", "validation-error", err.Error())
		return entities.Transaction{}, err
	}

	// Check account exists for this transaction
	_, err = repo.AccountsDatastore.GetAccount(ctx, input.AccountID)
	if err != nil {
		// Check if err is because the account does not exist vs some other issue
		if _, ok := err.(*entities.AccountNotFound); ok {
			// Account does not exist
			repo.Logger.Error("Cannot create transaction as account does not exist", "account-id", input.AccountID)
		} else {
			repo.Logger.Error("failed to check for account in datastore when creating transaction", "datastore-err", err.Error())
		}
		return entities.Transaction{}, err
	}

	// Create new uuid
	id, err := tc.uuidGen.New()
	if err != nil {
		repo.Logger.Error("failed to generate uuid for new transaction", "uuid-gen-err", err.Error())
		return entities.Transaction{}, err
	}

	now := tc.clock.Now()

	// Set amount based on operation type
	if input.OperationID != int(entities.Payment) {
		// amount should be recorded as negative
		input.Amount = -input.Amount
	}

	newTransaction := entities.Transaction{
		ID:            id,
		AccountID:     input.AccountID,
		OperationType: entities.OperationType(input.OperationID),
		Amount:        input.Amount,
		EventDate:     now,
	}

	err = repo.TransactionsDatastore.CreateTransaction(ctx, newTransaction)
	if err != nil {
		repo.Logger.Error("Failed to create transaction in database", "datatore-err", err.Error())
		return entities.Transaction{}, err
	}

	return newTransaction, nil

}
