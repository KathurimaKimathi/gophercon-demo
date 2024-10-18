package mock

import (
	"context"

	"github.com/KathurimaKimathi/gophercon-demo/pkg/gophercon/domain"
	"github.com/brianvoe/gofakeit"
)

var (
	id = gofakeit.UUID()
)

type DataStoreMock struct {
	MockCreateUserFn            func(ctx context.Context, user *domain.User) (*domain.User, error)
	MockCreateBusinessProfileFn func(ctx context.Context, user *domain.BusinessProfile) (*domain.BusinessProfile, error)
	MockCreateBusinessPartnerFn func(ctx context.Context, user *domain.BusinessPartner) (*domain.BusinessPartner, error)
}

// NewDataStoreMock returns a new instance of the mock datastore
func NewDataStoreMock() *DataStoreMock {
	return &DataStoreMock{
		MockCreateUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			return &domain.User{
				ID:        &id,
				Username:  gofakeit.BeerName(),
				FirstName: gofakeit.BeerName(),
				LastName:  gofakeit.BeerName(),
				Email:     gofakeit.Email(),
				UserType:  gofakeit.BeerName(),
			}, nil
		},
		MockCreateBusinessProfileFn: func(ctx context.Context, user *domain.BusinessProfile) (*domain.BusinessProfile, error) {
			return &domain.BusinessProfile{
				ID: &id,
			}, nil
		},
		MockCreateBusinessPartnerFn: func(ctx context.Context, user *domain.BusinessPartner) (*domain.BusinessPartner, error) {
			return &domain.BusinessPartner{
				ID: &id,
			}, nil
		},
	}
}

// CreateUser mocks the implementation of creating a new user
func (m *DataStoreMock) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return m.MockCreateUserFn(ctx, user)
}

// CreateBusinessProfile mocks the implementation of creating a new business profile
func (m *DataStoreMock) CreateBusinessProfile(ctx context.Context, user *domain.BusinessProfile) (*domain.BusinessProfile, error) {
	return m.MockCreateBusinessProfileFn(ctx, user)
}

// CreateBusinessPartner mocks the implementation of creating a new business partner
func (m *DataStoreMock) CreateBusinessPartner(ctx context.Context, user *domain.BusinessPartner) (*domain.BusinessPartner, error) {
	return m.MockCreateBusinessPartnerFn(ctx, user)
}
