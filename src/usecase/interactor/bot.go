package interactor

import (
	"context"
	"math"

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
		result := longestCommonSubsequence(v.Tag(), term)
		if len(result) > 4 {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func longestCommonSubsequence(str1, str2 string) string {
	table := make([][]int, len(str1)+1)
	for i := range table {
		table[i] = make([]int, len(str2)+1)
	}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				table[i+1][j+1] = table[i][j] + 1
			} else {
				table[i+1][j+1] = int(math.Max(float64(table[i+1][j]), float64(table[i][j+1])))
			}
		}
	}
	var result string
	for x, y := len(str1), len(str2); x != 0 && y != 0; {
		if table[x][y] == table[x-1][y] {
			x--
		} else if table[x][y] == table[x][y-1] {
			y--
		} else {
			result = string(str1[x-1]) + result
			x--
			y--
		}
	}
	return result
	
}
