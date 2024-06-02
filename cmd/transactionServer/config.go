package main

import (
	"io"
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	REST ServerConfig
}

type ServerConfig struct {
	Address string
	Port int
}

func LoadConfigFromFile(tomlPath string) *Config {
	// Open toml file
	file, err := os.Open(tomlPath)
	if err != nil {
		log.Fatal("Failed to open toml config file", err.Error())
	}
	defer file.Close()

	// Read toml file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read from  toml config file", err.Error())
	}

	// Unmarshall into config
	var config Config
	err = toml.Unmarshal(fileBytes, &config)
	if err != nil {
		log.Fatal("Failed to Unmarshal config toml", err.Error())
	}

	return &config
}
