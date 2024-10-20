package psql

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	log "github.com/sirupsen/logrus"

	// responsible for providing methods that are by gorm to connect to cloud SQL
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (

	// For gcp cloud postgres database
	GoogleProject     = "GOOGLE_CLOUD_PROJECT_ID"
	DatabaseRegion    = "DATABASE_REGION"
	DatabasesInstance = "DATABASE_INSTANCE"

	DBHost     = "POSTGRES_HOST"
	DBPort     = "POSTGRES_PORT"
	DBUser     = "POSTGRES_USER"
	DBPassword = "POSTGRES_PASSWORD"
	DBName     = "POSTGRES_DB"
)

type connectionConfig struct {
	host            string
	port            string
	user            string
	password        string
	dbname          string
	project         string
	region          string
	instance        string
	asCloudInstance bool
}

// PGInstance box for postgres client. We use this instead of a global variable
type PGInstance struct {
	DB *gorm.DB
}

// NewPGInstance creates a new instance of postgres client
func NewPGInstance() (*PGInstance, error) {
	db := startDatabase()
	if db == nil {
		return nil, fmt.Errorf("failed to start database: %v", db)
	}

	return &PGInstance{
		DB: db,
	}, nil
}

// isLocalDB returns true if the service is currently configured to use a local
// database.
func isLocalDB() bool {
	isLocal, err := strconv.ParseBool(os.Getenv("IS_LOCAL_DB"))
	if err != nil {
		return false
	}

	return isLocal
}

// startDatabase boots up the database
func startDatabase() *gorm.DB {
	user := helpers.MustGetEnvVar(DBUser)
	dbpassword := helpers.MustGetEnvVar(DBPassword)
	dbname := helpers.MustGetEnvVar(DBName)

	var config connectionConfig
	if isLocalDB() {
		config.host = helpers.MustGetEnvVar(DBHost)
		config.port = helpers.MustGetEnvVar(DBPort)
		config.user = user
		config.password = dbpassword
		config.dbname = dbname
	} else {
		config.project = helpers.MustGetEnvVar(GoogleProject)
		config.region = helpers.MustGetEnvVar(DatabaseRegion)
		config.instance = helpers.MustGetEnvVar(DatabasesInstance)
		config.asCloudInstance = true
		config.user = user
		config.password = dbpassword
		config.dbname = dbname
	}

	return boot(config)
}

func boot(cfg connectionConfig) *gorm.DB {
	var (
		err error
		db  *gorm.DB
	)

	if cfg.asCloudInstance {
		connString := fmt.Sprintf("host=%v:%v:%v user=%v dbname=%v password=%v sslmode=disable",
			cfg.project, cfg.region, cfg.instance, cfg.user, cfg.dbname, cfg.password)
		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        connString,
		}), &gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	} else {
		// called when using localhost instance of postgres
		connString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbname)
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if db == nil || err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Check that we can connect to the DB. This allows us to know early on if
	// we have the correct credentials and connection settings. As a bonus, we
	// also get a more descriptive error message and exit code.
	if connection, err := db.DB(); err != nil {
		log.Errorf("unable to get a connection to the database: %s", err)
		os.Exit(1)
	} else {
		if err = connection.Ping(); err != nil {
			log.Errorf("unable to ping the database: %s", err)
			os.Exit(1)
		}
	}

	return db
}
