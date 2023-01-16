package main

import (
	"os"

	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
	"github.com/joaquinicolas/iris-bot/src/registry"
)

func main() {

	gsheet := datastore.NewGsheetProvider(os.Getenv("GSHEET_TOKEN"))
	ch := make(chan bool)
	container := registry.NewRegistry(gsheet)
	container.Bootstrap()
	<-ch
}
