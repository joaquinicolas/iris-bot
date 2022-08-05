package repository

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
)

type BotRepository interface {
	GetProducts(ctx context.Context, term string) ([]entities.Product, error)
}


