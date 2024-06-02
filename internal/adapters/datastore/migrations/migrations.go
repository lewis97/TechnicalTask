package migrations

import (
	
)

import (
	"embed"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed *.sql
var FS embed.FS

func GetMigrations() *migrate.EmbedFileSystemMigrationSource {
	return &migrate.EmbedFileSystemMigrationSource{
		FileSystem: FS,
		Root:       ".",
	}
}
