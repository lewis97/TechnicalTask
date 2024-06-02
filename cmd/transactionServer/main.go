package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/lewis97/TechnicalTask/internal/adapters/datastore"
	"github.com/lewis97/TechnicalTask/internal/adapters/datastore/migrations"
	"github.com/lewis97/TechnicalTask/internal/adapters/server"
	"github.com/lewis97/TechnicalTask/internal/drivers/postgres"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {

	// Load config
	configPath := os.Getenv("CONFIG_FILE")
	if len(configPath) == 0 {
		log.Fatal("No config path specified in env variable CONFIG_FILE")
	}
	config := LoadConfigFromFile(configPath)

	// Connect to database
	db := postgres.NewDBConnection(config.Database)
	defer db.Close()
	ds := datastore.NewDatastore(db)

	// Run DB migrations
	migration := migrations.GetMigrations()
	migrate.SetTable("migrations")
	_, err := migrate.Exec(db.DB, "postgres", migration, migrate.Up)
	if err != nil {
		log.Fatal("DB Migrations failed: ", err.Error())
	}

	// Set logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup & start REST server
	deps := server.Dependencies{
		Logger: *logger,
		Datastore: ds,
	}
	server := server.New(deps)
	server.Start(config.REST.Address, config.REST.Port)
	// TODO: ^ goroutine w/ context?
}
