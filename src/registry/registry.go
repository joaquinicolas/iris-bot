package registry

import "github.com/joaquinicolas/iris-bot/src/infrastructure/datastore"

type registry struct {
	store *datastore.Store
}

type Registry interface {
	
}

func NewRegistry(store *datastore.Store) Registry {
	return &registry{}
}

