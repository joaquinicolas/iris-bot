package presenter

import (
	"context"
	"fmt"
	"io"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
)

type BotPresenter interface {
	RespondWithProducts(products []entities.Product) (string, int)
}

type botPresenter struct {
	w   io.Writer
	ctx context.Context
}

// RespondWithProducts implements BotPresenter
func (*botPresenter) RespondWithProducts(products []entities.Product) (string, int) {
	// concatenate the products into a string
	var productsString string
	var counter int
	for _, product := range products {
		counter++
		productsString += fmt.Sprintf("%d. %s:%s\n", counter, product.Name, product.Price)
	}

	return productsString, counter
}

// RespondWithProducts implements io.Writer
func (bp *botPresenter) Write(p []byte) (n int, err error) {
	bp.w.Write(p)
	return len(p), nil
}

// method to create a new instance of the presenter
func NewBotPresenter(ctx context.Context) BotPresenter {
	return &botPresenter{ctx: ctx}
}
