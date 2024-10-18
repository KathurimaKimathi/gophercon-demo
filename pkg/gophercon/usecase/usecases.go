package usecase

import (
	"context"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/infrastructure"
)

type IUsecase interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

// Gophercon models the configuration needed by the usecase
type Gophercon struct {
	Infrastructure infrastructure.Infrastructure
}

// NewUseCasesInteractor initializes a new usecases interactor
func NewUseCasesInitializer(
	infra infrastructure.Infrastructure,
) IUsecase {
	return &Gophercon{
		Infrastructure: infra,
	}
}
