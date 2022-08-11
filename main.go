package main

import (
	"os"

	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
	"github.com/joaquinicolas/iris-bot/src/registry"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	gsheet := datastore.NewGsheetProvider(os.Getenv("GSHEET_TOKEN"))
	ch := make(chan bool)
	container := registry.NewRegistry(gsheet)
	container.Bootstrap()
	<-ch
}
