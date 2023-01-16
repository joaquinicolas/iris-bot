package presenter

import (
	"context"
	"fmt"
	"io"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/joaquinicolas/iris-bot/src/usecase/presenter"
)

type botPresenter struct {
	w   io.Writer
	ctx context.Context
}

// RespondWithProducts implements BotPresenter
func (*botPresenter) RespondWithProducts(products []entities.Product) (string, int) {
	// concatenate the products into a string
	var productsString string = `Se han encontrado los siguientes productos:`
	var counter int
	for _, product := range products {
		counter++
		productsString += fmt.Sprintf("\n<b>%d.</b> %s: <b>%s</b>", counter, product.Name, product.Price)
	}

	return productsString, counter
}

// RespondWithProducts implements io.Writer
func (bp *botPresenter) Write(p []byte) (n int, err error) {
	return bp.w.Write(p)
}

// method to create a new instance of the presenter
func NewBotPresenter(ctx context.Context) presenter.BotPresenter {
	return &botPresenter{ctx: ctx}
}
