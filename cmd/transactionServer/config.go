package main

import (
	"io"
	"log"
	"os"

	"github.com/lewis97/TechnicalTask/internal/drivers/postgres"
	toml "github.com/pelletier/go-toml"
)

// Config contains the REST and Database configuration options needed for the app.
// These are read in via a toml file referenced in the environment variable CONFIG_FILE
// When running locally these will be stored in bench-config.toml.
type Config struct {
	REST     ServerConfig
	Database postgres.DatabaseConfig
}

type ServerConfig struct {
	Address string
	Port    int
}

func LoadConfigFromFile(tomlPath string) *Config {
	// Open toml file
	file, err := os.Open(tomlPath)
	if err != nil {
		log.Fatal("Failed to open toml config file", err.Error())
	}
	defer file.Close()

	// Read toml file into memory
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read from  toml config file", err.Error())
	}

	// Unmarshal into config struct
	var config Config
	err = toml.Unmarshal(fileBytes, &config)
	if err != nil {
		log.Fatal("Failed to Unmarshal config toml", err.Error())
	}

	return &config
}
