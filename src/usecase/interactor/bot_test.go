package interactor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetProducts(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Test_GetProducts ok", wantErr: false},
		{name: "Test_GetProducts get error", wantErr: true},
	}

	bot := NewBotInteractor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products, err := bot.GetProducts(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (len(products) == 0) && !tt.wantErr {
				t.Errorf("GetProducts() got empty products")
			}

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
		name    string
		term    string
		wantErr bool
	}{
		{name: "Test_GetProducts ok", wantErr: false, term: "queso"},
		{name: "Test_GetProducts get unexpected error", wantErr: true, term:""},
		{name: "Test_GetProducts get error", wantErr: true, term: "error"},
	}

	bot := NewBotInteractor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products, err := bot.GetProductsByTerm(context.Background(), "")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (len(products) == 0) && !tt.wantErr {
				t.Errorf("GetProducts() got empty products")
			}

			for _, v := range products {
				assert.NotNil(t, v.Id)
				assert.NotNil(t, v.Name)
				assert.NotNil(t, v.Price)
			}
		})
	}
}
