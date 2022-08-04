package interactor

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
)

type botInteractor struct {

}

type BotInteractor interface {
	GetProducts(ctx context.Context) ([]entities.Product, error)
	GetProductsByTerm(ctx context.Context, term string) ([]entities.Product, error)
}

func NewBotInteractor(
	
) BotInteractor {
	return &botInteractor{}
}


func (bi *botInteractor) GetProducts(ctx context.Context) ([]entities.Product, error) {
	return nil, nil
}

func (bi *botInteractor) GetProductsByTerm(ctx context.Context, term string) ([]entities.Product, error) {
	return nil, nil
}
