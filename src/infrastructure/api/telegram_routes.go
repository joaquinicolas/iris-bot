package api

import (
	"context"
	"fmt"

	objs "github.com/SakoDroid/telego/objects"
	"github.com/joaquinicolas/iris-bot/src/usecase/interactor"
	"github.com/joaquinicolas/iris-bot/src/usecase/presenter"
)

type TelegramRouter struct {
	tb                *TelegramBot
	botInteractor     interactor.BotInteractor
	telegramPresenter presenter.BotPresenter
}

func NewTelegramRouter(
	telegramBot *TelegramBot,
	interactor interactor.BotInteractor,
	presenter presenter.BotPresenter,
) TelegramRouter {
	return TelegramRouter{
		tb:                telegramBot,
		botInteractor:     interactor,
		telegramPresenter: presenter,
	}
}

func (tr *TelegramRouter) Register() {
	// nolint
	tr.tb.bot.AddHandler("(h|H)ola", func(u *objs.Update) {
		tr.start(u)
	}, "all")

	// nolint
	tr.tb.bot.AddHandler(".*", func(u *objs.Update) {
		tr.getProductsByTerm(u)
	}, "all")
}

func (tr *TelegramRouter) start(u *objs.Update) {
	_, err := tr.tb.bot.SendMessage(
		u.Message.Chat.Id,
		"Hola, mi nombre es Iris, para ayudarte necesito que me digas un producto: ",
		"",
		u.Message.MessageId,
		false,
		false)
	if err != nil {
		fmt.Println(err)
	}
}

// tr *TelegramRouter getProductsByTerm
func (tr *TelegramRouter) getProductsByTerm(u *objs.Update) {
	ctx := context.Background()
	products, err := tr.botInteractor.GetProductsByTerm(ctx, u.Message.Text)
	if err != nil {
		fmt.Println(err)
	}
	text, _ := tr.telegramPresenter.RespondWithProducts(products)
	_, err = tr.tb.bot.SendMessage(u.Message.Chat.Id, text, "html", u.Message.MessageId, false, false)
	if err != nil {
		fmt.Println(err)
	}
}

