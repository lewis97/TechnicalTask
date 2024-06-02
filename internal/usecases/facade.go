package usecases

import (
	"context"

	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
)


type Facade struct {
	AccountsUsecase *accounts.AccountsUsecase
}

func (f *Facade) GetAccount(ctx context.Context, input *accounts.GetAcccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error){
	return f.AccountsUsecase.GetAccount(ctx, input, repo)
}

func (f *Facade) CreateAccount(ctx context.Context, input *accounts.CreateAccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error) {
	return f.AccountsUsecase.CreateAccount(ctx, input, repo)
}

