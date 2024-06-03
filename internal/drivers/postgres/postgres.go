package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

// Establish and return a new database connection

func NewDBConnection(config DatabaseConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := sqlx.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err.Error())
	}

	return db
}
