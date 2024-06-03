package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/lewis97/TechnicalTask/internal/adapters/datastore"
	"github.com/lewis97/TechnicalTask/internal/adapters/datastore/migrations"
	"github.com/lewis97/TechnicalTask/internal/adapters/server"
	"github.com/lewis97/TechnicalTask/internal/drivers/postgres"
	"github.com/lewis97/TechnicalTask/internal/drivers/uuidgen"
	"github.com/lewis97/TechnicalTask/internal/drivers/clock"
	"github.com/lewis97/TechnicalTask/internal/usecases"
	"github.com/lewis97/TechnicalTask/internal/usecases/accounts"
	"github.com/lewis97/TechnicalTask/internal/usecases/transactions"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {

	// Load config
	configPath := os.Getenv("CONFIG_FILE")
	if len(configPath) == 0 {
		log.Fatal("No config path specified in env variable CONFIG_FILE")
	}
	config := LoadConfigFromFile(configPath)

	// Set logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Connect to database
	db := postgres.NewDBConnection(config.Database)
	defer db.Close()
	ds := datastore.NewDatastore(db, *logger)

	// Run DB migrations
	migration := migrations.GetMigrations()
	migrate.SetTable("migrations")
	_, err := migrate.Exec(db.DB, "postgres", migration, migrate.Up)
	if err != nil {
		log.Fatal("DB Migrations failed: ", err.Error())
	}

	uuidGenerator := uuidgen.NewGoogleUUIDGen()
	clock := clock.NewTimeClock()

	// Setup & start REST server
	deps := server.Dependencies{
		Logger:    *logger,
		Datastore: ds,
		Usecases: &usecases.Facade{
			AccountsUsecase:     accounts.NewAccountsUsecase(uuidGenerator, clock),
			TransactionsUsecase: transactions.NewTransactionUsecase(uuidGenerator, clock),
		},
	}
	server := server.New(deps)
	server.Start(config.REST.Address, config.REST.Port)
	// TODO: ^ goroutine w/ context?
}
