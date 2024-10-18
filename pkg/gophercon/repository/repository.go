package repository

import (
	"context"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
)

type IUser interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
