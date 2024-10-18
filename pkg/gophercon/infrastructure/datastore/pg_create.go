package datastore

import (
	"context"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure/datastore/psql"
	"github.com/mitchellh/mapstructure"
)

// CreateUser maps the user's domain data to user model data
func (d *DbServiceImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	usr := &psql.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserType:  user.UserType,
	}

	result, err := d.create.CreateUser(ctx, usr)
	if err != nil {
		return nil, err
	}

	var output *domain.User
	if err := mapstructure.Decode(result, &output); err != nil {
		return nil, err
	}

	return output, nil
}
