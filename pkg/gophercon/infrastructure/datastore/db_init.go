package datastore

import (
	"context"
	"log"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/application/common/helpers"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure/datastore/psql"
)

type DBCreate interface {
	CreateUser(ctx context.Context, user *psql.User) (*psql.User, error)
}

type DBQuery interface {
}

type DBUpdate interface {
}

// DbServiceImpl is an implementation of the database repository
// It is implementation agnostic i.e logic should be handled using
// the preferred database
type DbServiceImpl struct {
	create DBCreate
	query  DBQuery
	update DBUpdate
}

// NewDbService creates a new database service
func NewDbService() *DbServiceImpl {
	// This implementation is database agnostic. It can be changed to use any database. e.g. Firebase, MongoDB, etc
	environment := helpers.MustGetEnvVar("REPOSITORY")

	switch environment {
	case "firebase":
		return &DbServiceImpl{}

	case "postgres":
		pg, err := psql.NewPGInstance()
		if err != nil {
			log.Panicf("can't initialize postgres when setting up profile service: %s", err)
		}

		return &DbServiceImpl{
			create: pg,
			query:  pg,
			update: pg,
		}

	default:
		log.Panicf("unknown repository: %s", environment)
	}

	return nil
}
