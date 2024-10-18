package migration_asset

import "embed"

//go:embed migrations/*.sql
var DBMigrations embed.FS
