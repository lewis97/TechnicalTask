package main

import (
	"log"
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

	// Setup & start REST server
	server := server.New(server.Dependencies{})
	server.Start(config.REST.Address, config.REST.Port)
}
