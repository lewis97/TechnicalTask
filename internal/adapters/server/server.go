package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/lewis97/TechnicalTask/internal/adapters/datastore"
	"github.com/lewis97/TechnicalTask/internal/domain/entities"
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
	"github.com/lewis97/TechnicalTask/internal/usecases/transactions"
)

type Usecase interface {
	GetAccount(ctx context.Context, input *accounts.GetAcccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error)
	CreateAccount(ctx context.Context, input *accounts.CreateAccountInput, repo *accounts.AccountUsecaseRepos) (entities.Account, error)
	CreateTransaction(ctx context.Context, input *transactions.CreateTransactionInput, repo *transactions.TransactionsUsecaseRepos) (entities.Transaction, error)
}

type Dependencies struct {
	Logger    slog.Logger
	Datastore *datastore.Datastore
	Usecases  Usecase
}

type Server struct {
	routes    *http.ServeMux
	logger    slog.Logger
	datastore *datastore.Datastore
	usecases  Usecase
}

func (s Server) Start(addr string, port int) {
	// log.Println(fmt.Sprintf("server started. listening on port %d", port))
	s.logger.Info(
		"Server started and listening",
		"address", addr,
		"port", port,
	)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), s.routes); err != nil {
		log.Fatalln("server startup failed, %w", err)
		// TODO: Replace with normal log & pass this error back to main.go to the log.fatal
	}
}

func New(deps Dependencies) *Server {
	router := http.NewServeMux()
	apiConfig := huma.DefaultConfig("Transaction service REST API", "1.0.0")
	// Modify the config to include schema in response
	apiConfig.CreateHooks = []func(huma.Config) huma.Config{}
	api := humago.New(router, apiConfig)

	s := &Server{
		routes:    router,
		logger:    deps.Logger,
		datastore: deps.Datastore,
		usecases:  deps.Usecases,
	}

	// Map routes
	huma.Get(api, "/hello", s.TestHandler)
	huma.Get(api, "/accounts/{accountID}", s.GetAccount) // get account by ID
	huma.Post(api, "/accounts", s.CreateAccount)         // create account
	huma.Post(api, "/transactions", s.CreateTransaction) // create transaction

	return s
}
