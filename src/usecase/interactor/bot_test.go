package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/joaquinicolas/iris-bot/src/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetProducts(t *testing.T) {
	tests := []struct {
		name     string
		expected []entities.Product
		wantErr  bool
	}{
		{name: "Test_GetProducts ok", wantErr: false, expected: mocks.MockOkResponse()},
		{name: "Test_GetProducts get error", wantErr: true, expected: nil},
	}

	repo := mocks.NewRepository()
	bot := NewBotInteractor(
		repo,
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if tt.wantErr {
				repo.On("GetProducts", ctx).Return(tt.expected, errors.New("got an error")).Once()
			} else {
				repo.On("GetProducts", ctx).Return(tt.expected, nil).Once()
			}
			products, err := bot.GetProducts(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.ObjectsAreEqualValues(tt.expected, products)
			for _, v := range products {
				assert.NotNil(t, v.Id)
				assert.NotNil(t, v.Name)
				assert.NotNil(t, v.Price)
			}
		})
	}
}

func Test_GetProductBySearchTerm(t *testing.T) {
	tests := []struct {
		name     string
		expected []entities.Product
		term     string
		wantErr  bool
	}{
		{name: "Test_GetProducts ok", wantErr: false, term: "test", expected: []entities.Product{{Id: "4", Name: "test 70GR", Price: "4"}}},
		{name: "Test_GetProducts get an error", wantErr: true, term: "", expected: nil},
		{name: "Test_GetProducts get unexpected error", wantErr: true, term: "error", expected: nil},
	}

	repo := mocks.NewRepository()
	bot := NewBotInteractor(
		repo,
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if tt.wantErr {
				repo.On("GetProducts", ctx).Return(tt.expected, errors.New("got an error")).Once()
			} else {
				repo.On("GetProducts", ctx).Return(tt.expected, nil).Once()
			}
			products, err := bot.GetProductsByTerm(ctx, tt.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.ObjectsAreEqualValues(tt.expected, products)
			for _, v := range products {
				assert.NotNil(t, v.Id)
				assert.NotNil(t, v.Name)
				assert.NotNil(t, v.Price)
			}
		})
	}
}
