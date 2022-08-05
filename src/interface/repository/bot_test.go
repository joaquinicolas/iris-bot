package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/joaquinicolas/iris-bot/src/interface/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_botRepository_GetProducts(t *testing.T) {
	tests := []struct {
		name     string
		term     string
		expected []entities.Product
		wantErr  bool
	}{
		{
			name:     "Get an unexpected error",
			term:     "",
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Get a list of products filtered by test",
			term: "test",
			expected: []entities.Product{
				{Id: "4", Name: "test 70GR", Price: "4"},
			},
			wantErr: false,
		},
	}
	store := new(mocks.Store)
	bot := NewBotRepository(store)
	for _, tt := range tests {
		tt := tt
		if tt.wantErr {
			store.On("Get", "", "").Return([]entities.Product{}, errors.New("unexpected error")).Once()
		} else {
			store.On("Get", "", "").Return(mocks.MockOkResponse(), nil).Once()
		}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			products, err := bot.GetProducts(ctx, tt.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.True(t, assert.ObjectsAreEqualValues(tt.expected, products))
			for _, v := range products {
				assert.NotEmpty(t, v.Id)
				assert.NotEmpty(t, v.Name)
				assert.NotEmpty(t, v.Price)
			}
		})
	}
}
