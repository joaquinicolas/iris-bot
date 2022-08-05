package repository

import (
	"context"
	"fmt"
	"regexp"

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

func (br *botRepository) GetProducts(ctx context.Context, term string) ([]entities.Product, error) {
	if err := br.store.Connect(ctx); err != nil {
		return nil, err
	}

	readRange, sheetId := getConfig()
	products, err := br.store.Get(sheetId, readRange)
	if err != nil {
		return nil, err
	}

	if term == "" {
		return products, nil
	} else {
		return filterByTerm(products, term), nil
	}
}

func getConfig() (string, string) {
	return "", ""
}

func filterByTerm[T entities.Filter](list []T, term string) []T {
	r, _ := regexp.Compile(fmt.Sprintf("[a-zA-Z]*\\s*%s\\s*[a-zA-Z]*", term))
	filtered := make([]T, 0)
	for _, v := range list {
		if r.MatchString(v.Tag()) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
