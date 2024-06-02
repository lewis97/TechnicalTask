package server

import (
	"fmt"
	"log"
	"net/http"
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)


type Dependencies struct {
	Logger slog.Logger
}

type Server struct {
	routes         *http.ServeMux
	logger 		   slog.Logger
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
		routes:         router,
		logger: deps.Logger,

	}

	huma.Get(api, "/hello", s.TestHandler)

	return s
}
