package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/lewis97/TechnicalTask/internal/adapters/server"
)

func main() {

	// Load config
	configPath := os.Getenv("CONFIG_FILE")
	if len(configPath) == 0 {
		log.Fatal("No config path specified in env variable CONFIG_FILE")
	}
	config := LoadConfigFromFile(configPath)
	log.Println(config.REST.Address)

	// Set logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup & start REST server
	deps := server.Dependencies{
		Logger: *logger,
	}
	server := server.New(deps)
	server.Start(config.REST.Address, config.REST.Port)
	// TODO: ^ goroutine w/ context?
}
