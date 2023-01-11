package mocks

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

type Store struct {
	mock.Mock
}

func (s *Store) Connect(ctx context.Context) error {
	return nil
}

func (s *Store) Get(sheetId string, readRange string) ([]entities.Product, error) {
	args := s.Called(sheetId, readRange)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func MockOkResponse() []entities.Product {
	return []entities.Product{
		{Id: "1", Name: "pepitos", Price: "1"},
		{Id: "2", Name: "oreo", Price: "2"},
		{Id: "3", Name: "chocolate 35GR", Price: "3"},
		{Id: "4", Name: "test 70GR", Price: "4"},
		{Id: "4", Name: "test", Price: "2"},
		{Id: "4", Name: "test 80GR", Price: "5"},
	}
}
