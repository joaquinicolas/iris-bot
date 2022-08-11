package interactor

import (
	"context"
	"fmt"
	"regexp"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/joaquinicolas/iris-bot/src/usecase/repository"
)

type botInteractor struct {
	repository repository.BotRepository
}

type BotInteractor interface {
	GetProducts(ctx context.Context) ([]entities.Product, error)
	GetProductsByTerm(ctx context.Context, term string) ([]entities.Product, error)
}

func NewBotInteractor(repository repository.BotRepository) BotInteractor {
	return &botInteractor{
		repository: repository,
	}
}

func (bi *botInteractor) GetProducts(ctx context.Context) ([]entities.Product, error) {
	products, err := bi.repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (bi *botInteractor) GetProductsByTerm(ctx context.Context, term string) ([]entities.Product, error) {
	products, err := bi.repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	if term == "" {
		return products, nil
	}
	return filterByTerm(products, term), nil
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
