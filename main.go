package main

import (
	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
	"github.com/joaquinicolas/iris-bot/src/registry"
)

const (
	gsheetToken = "AIzaSyB05LD8IrWbvmQLR6a0dUicsoePnGirsH8"
	telegramToken = "5473899125:AAHRn8jEWfjk1vNJDuwYL9AiKWRUTHXGeBM"
	sheetId = "1x3UNLdLbmnl0d65fEmH2us4Xj15Xj9ZQ-UNM5SXFVDg"
	sheetRange = "PRECIOS TOTALES!A3:O90"
)

func main() {
	gsheet := datastore.NewGsheetProvider(gsheetToken)
	ch := make(chan bool)
	container := registry.NewRegistry(gsheet)
	container.Bootstrap()
	<-ch
}
