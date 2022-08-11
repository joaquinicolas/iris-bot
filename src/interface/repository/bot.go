package repository

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
)

type botRepository struct {
	store datastore.Store
}

func NewBotRepository(
	store datastore.Store,
) *botRepository {
	return &botRepository{
		store: store,
	}
}

func (br *botRepository) GetProducts(ctx context.Context) ([]entities.Product, error) {
	if err := br.store.Connect(ctx); err != nil {
		return nil, err
	}

	readRange, sheetId := getConfig()
	products, err := br.store.Get(sheetId, readRange)
	if err != nil {
		return nil, err
	}

	return products, nil

}

func getConfig() (string, string) {
	return "PRECIOS TOTALES!A3:O122", "1x3UNLdLbmnl0d65fEmH2us4Xj15Xj9ZQ-UNM5SXFVDg"
}
