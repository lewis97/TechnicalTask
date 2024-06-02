package usecases

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
	"github.com/lewis97/TechnicalTask/internal/usecases/transactions"
)


type Facade struct {
	AccountsUsecase *accounts.AccountsUsecase
	TransactionsUsecase *transactions.TransactionsUsecase
}

func (f *Facade) GetAccount(ctx context.Context, input *accounts.GetAcccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error){
	return f.AccountsUsecase.GetAccount(ctx, input, repo)
}

func (f *Facade) CreateAccount(ctx context.Context, input *accounts.CreateAccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error) {
	return f.AccountsUsecase.CreateAccount(ctx, input, repo)
}

func (f *Facade) CreateTransaction(ctx context.Context, input *transactions.CreateTransactionInput, repo *transactions.TransactionsUsecaseRepos) (entities.Transaction, error){
	return f.TransactionsUsecase.CreateTransaction(ctx, input, repo)
}
