package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/lewis97/TechnicalTask/internal/adapters/datastore"
)

type Dependencies struct {
	Logger    slog.Logger
	Datastore *datastore.Datastore
}

type Server struct {
	routes    *http.ServeMux
	logger    slog.Logger
	datastore *datastore.Datastore
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
	api := humago.New(router, huma.DefaultConfig("Transaction service REST API", "1.0.0"))

	s := &Server{
		routes:    router,
		logger:    deps.Logger,
		datastore: deps.Datastore,
	}

	// Map routes
	huma.Get(api, "/hello", s.TestHandler)
	huma.Get(api, "/accounts/{accountID}", s.GetAccount) // get account by ID
	huma.Post(api, "/accounts", s.CreateAccount)         // create account
	huma.Post(api, "/transactions", s.CreateTransaction) // create transaction

	return s
}
