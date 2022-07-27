package main

import (
	"context"

	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
)

func main() {
	/*token := "5473899125:AAHRn8jEWfjk1vNJDuwYL9AiKWRUTHXGeBM"
	bot := api.NewTelegramBot(token)
	err := bot.Run()
	ch := make(chan bool)
	if err != nil {
		panic(err)
	}

	api.NewRouter(bot)
	<-ch

	1x3UNLdLbmnl0d65fEmH2us4Xj15Xj9ZQ-UNM5SXFVDg

	*/

	gsheet := datastore.NewGsheetProvider("AIzaSyB05LD8IrWbvmQLR6a0dUicsoePnGirsH8")
	err := gsheet.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	gsheet.Get("1x3UNLdLbmnl0d65fEmH2us4Xj15Xj9ZQ-UNM5SXFVDg")
}
