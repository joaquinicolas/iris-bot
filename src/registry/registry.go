package registry

import (
	"context"
	"os"

	"github.com/joaquinicolas/iris-bot/src/infrastructure/api"
	"github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"
	"github.com/joaquinicolas/iris-bot/src/interface/presenter"
	"github.com/joaquinicolas/iris-bot/src/interface/repository"
	"github.com/joaquinicolas/iris-bot/src/usecase/interactor"
)

type registry struct {
	store datastore.Store
}

type Registry interface {
	Bootstrap() 
}

// registry implements the Registry interface
func (r *registry) Bootstrap() {
	botRepository := repository.NewBotRepository(r.store)
	service := interactor.NewBotInteractor(botRepository)
	telegramBot := api.NewTelegramBot(os.Getenv("TELEGRAM_TOKEN"))
	presenter := presenter.NewBotPresenter(context.Background())
	telegramRouter := api.NewTelegramRouter(telegramBot, service, presenter)
	err := telegramBot.Run()
	// TODO: hanldle error
	if err != nil {
		panic(err)
	}
	telegramRouter.Register()
}

func NewRegistry(store datastore.Store) Registry {
	return &registry{
		store: store,
	}
}
