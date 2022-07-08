package api

import (
	"fmt"
	objs "github.com/SakoDroid/telego/objects"
)

func NewRouter(tb *TelegramBot)  {
	tb.bot.AddHandler("/start", func(u *objs.Update) {
		_, err := tb.bot.SendMessage(u.Message.Chat.Id, "Hola, Mi nombre es Iris. En que puedo ayudarte?", "", u.Message.MessageId, false, false)
		if err != nil {
			fmt.Println(err)
		}
	}, "all")
}
