package interactor

import (
	"context"

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
	return filterByTerm(products, term), nil
}

func filterByTerm[T entities.Filter](list []T, term string) []T {
	filtered := make([]T, 0)
	for _, v := range list {
		result := lcsubStr(v.Tag(), term)
		if len(result) > 3 {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func lcsubStr(a, b string) string {
	m := len(a)
	n := len(b)

	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	var longestStr = ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1
				if len(longestStr) < table[i][j] {
					longestStr = a[i-table[i][j] : i+1]
				}
			} else {
				table[i][j] = 0
			}
		}
	}

	return longestStr
}
