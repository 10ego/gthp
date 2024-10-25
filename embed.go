package root

import "embed"

//go:embed internal/database/schema/*.sql
var embedMigrations embed.FS

func GetMigrationFS() embed.FS {
	return embedMigrations
}
