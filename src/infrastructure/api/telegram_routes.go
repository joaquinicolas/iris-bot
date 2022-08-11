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
	tr.tb.bot.AddHandler("(h|H)ola", func(u *objs.Update) {
		tr.start(u)
	}, "all")

	tr.tb.bot.AddHandler(".*", func(u *objs.Update) {
		tr.getProductsByTerm(u)
	}, "all")

	//go tr.inlineQuery()
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

func (tr *TelegramRouter) inlineQuery() {
	bot := tr.tb.bot
	presenter := tr.telegramPresenter
	inlineQueryChannel, _ := bot.AdvancedMode().RegisterChannel("", "inline_query")
	for in := range *inlineQueryChannel {

		iqs := bot.AdvancedMode().AAnswerInlineQuery(in.InlineQuery.Id, 0, false, "", "", "")
		products, err := tr.botInteractor.GetProductsByTerm(context.Background(), in.InlineQuery.Query)
		fmt.Println(products)
		fmt.Println(in.InlineQuery.Query)
		if err != nil {
			fmt.Println(err)
		}
		text, _ := presenter.RespondWithProducts(products)
		fmt.Println(text)
		_ = iqs.CreateTextMessage(
			"testest",
			"",
			nil,
			false,
		)

		_, err = iqs.Send()

		if err != nil {
			fmt.Println(err)
		}
	}
}
