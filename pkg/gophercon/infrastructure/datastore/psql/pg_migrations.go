package psql

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	// responsible for methods to connect to db for migrations to run
	migration_asset "github.com/KathurimaKimathi/gophercon-demo/db"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type migrationConfig struct {
	host     string
	port     string
	dbtype   string
	password string
	dbname   string
	user     string
}

// RunMigrations looks at the currently active migration version and will migrate all the way up (applying all up migrations)
func RunMigrations() error {
	driver, err := iofs.New(migration_asset.DBMigrations, "migrations")
	if err != nil {
		return err
	}

	migrationConfig := migrationConfig{
		dbtype:   "postgres",
		host:     helpers.MustGetEnvVar(DBHost),
		port:     helpers.MustGetEnvVar(DBPort),
		user:     helpers.MustGetEnvVar(DBUser),
		password: helpers.MustGetEnvVar(DBPassword),
		dbname:   helpers.MustGetEnvVar(DBName),
	}
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		migrationConfig.dbtype,
		migrationConfig.user,
		migrationConfig.password,
		migrationConfig.host,
		migrationConfig.port,
		migrationConfig.dbname,
	)

	m, err := migrate.NewWithSourceInstance("postgres", driver, connString)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			return nil
		}

		return err
	}

	return nil
}
