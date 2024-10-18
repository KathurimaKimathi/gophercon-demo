package usecase

import (
	"context"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
)

func (i *Gophercon) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// TODO: Implement some business logic here eg. verifying email existence
	return i.Infrastructure.Repository.CreateUser(ctx, user)
}
