package api

import (
	"fmt"

	objs "github.com/SakoDroid/telego/objects"
	"github.com/joaquinicolas/iris-bot/src/usecase/interactor"
)

type TelegramRouter struct {
	tb            *TelegramBot
	botInteractor interactor.BotInteractor
}

func NewTelegramRouter(telegramBot *TelegramBot, interactor interactor.BotInteractor) TelegramRouter {
	return TelegramRouter{
		tb:            telegramBot,
		botInteractor: interactor,
	}
}

func (tr *TelegramRouter) Register() {
	tr.tb.bot.AddHandler("/", func(u *objs.Update) {
		tr.start(u)
	}, "all")
}

func (tr *TelegramRouter) start(u *objs.Update) {
	_, err := tr.tb.bot.SendMessage(u.Message.Chat.Id, "Hola, Mi nombre es Iris. En que puedo ayudarte?", "", u.Message.MessageId, false, false)
	if err != nil {
		fmt.Println(err)
	}
}
