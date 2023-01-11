package mocks

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewRepository() *MockRepository  {
	return new(MockRepository)	
}

func (r *MockRepository) GetProducts(ctx context.Context) ([]entities.Product, error) {
	args := r.Called(ctx)
	return args.Get(0).([]entities.Product), args.Error(1)
}
