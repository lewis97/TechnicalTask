package postgres

// Config that holds all required config options for a database connection

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
