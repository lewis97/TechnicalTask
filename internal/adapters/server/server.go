package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)


type Dependencies struct {

}

type Server struct {
	routes         *http.ServeMux
}

func (s Server) Start(addr string, port int) {
	log.Println(fmt.Sprintf("server started. listening on port %d", port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), s.routes); err != nil {
		log.Fatalln("server startup failed, %w", err)
	}
}

func New(deps Dependencies) *Server {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Transaction service REST API", "1.0.0"))

	s := &Server{
		routes:         router,

	}

	huma.Get(api, "/hello", s.TestHandler)

	return s
}
