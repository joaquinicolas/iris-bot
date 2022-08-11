package presenter

import (
	"github.com/joaquinicolas/iris-bot/src/domain/entities"
)

type BotPresenter interface {
	RespondWithProducts(products []entities.Product) (string, int)
}

