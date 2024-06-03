package accounts

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/domain/repositories"
	"github.com/lewis97/TechnicalTask/internal/drivers/clock"
	"github.com/lewis97/TechnicalTask/internal/drivers/uuidgen"
)

// Accounts usecase

// Passing in a uuid generator to make mocking/UT easier
type AccountsUsecase struct {
	uuidGen uuidgen.UUIDGenerator
	clock   clock.Clock
}

type AccountUsecaseRepos struct {
	Logger            slog.Logger
	AccountsDatastore repositories.Accounts
}

func NewAccountsUsecase(uuidGenerator uuidgen.UUIDGenerator, clock clock.Clock) *AccountsUsecase {
	return &AccountsUsecase{
		uuidGen: uuidGenerator,
		clock:   clock,
	}
}

// Getting an account

type GetAcccountInput struct {
	AccountID uuid.UUID
}

func (ac *AccountsUsecase) GetAccount(ctx context.Context, input *GetAcccountInput, repo *AccountUsecaseRepos) (entities.Account, error) {
	account, err := repo.AccountsDatastore.GetAccount(ctx, input.AccountID)
	if err != nil {
		repo.Logger.Error(
			"Failed to get account from datastore",
			"accountID", input.AccountID.String(),
			"datastore-err", err.Error(),
		)
		return entities.Account{}, err
	}

	return *account, nil

}

// Creating an account

type CreateAccountInput struct {
	DocumentNumber uint
}

// Validate account creation input
func (input *CreateAccountInput) Validate() error {
	// No default values
	if input.DocumentNumber == 0 {
		return entities.NewInvalidInputError("Document number must be specified and cannot be 0")
	}
	return nil
}

func (ac *AccountsUsecase) CreateAccount(ctx context.Context, input *CreateAccountInput, repo *AccountUsecaseRepos) (entities.Account, error) {
	// Validate the input
	if validationErr := input.Validate(); validationErr != nil {
		repo.Logger.Error("validation of create account input failed", "validationError", validationErr.Error())
		return entities.Account{}, validationErr
	}

	// First check that we don't already have an account with that document number
	// Expect this to return an AccountNotFound
	_, err := repo.AccountsDatastore.GetAccountByDoc(ctx, input.DocumentNumber)
	if err == nil {
		// Account already exists with this doc ID
		repo.Logger.Error(
			"Cannot create account. Account already exists with documentID",
			"documentID",
			input.DocumentNumber)
		return entities.Account{}, entities.NewAccountAlreadyExistsError(input.DocumentNumber)
	} else if _, ok := err.(*entities.AccountNotFound); !ok {
		// Error is not an account not found error, therefore it's a genuine error
		repo.Logger.Error("Failed to check for existing account during account creation", "error", err.Error())
		return entities.Account{}, err
	}

	// Create a new id for the account
	id, err := ac.uuidGen.New()
	if err != nil {
		repo.Logger.Error("failed to generate uuid for new account", "uuid-gen-err", err.Error())
		return entities.Account{}, err
	}

	now := ac.clock.Now()

	newAccount := entities.Account{
		ID:             id,
		DocumentNumber: input.DocumentNumber,
		CreatedAt:      now,
	}

	// Create the account in the datastore
	err = repo.AccountsDatastore.CreateAccount(ctx, newAccount)
	if err != nil {
		repo.Logger.Error("Failed to commit new account to datastore", "datastore-error", err.Error())
		return entities.Account{}, err
	}

	repo.Logger.Info("successfully created new account", "accountID", id)
	return newAccount, nil

}
