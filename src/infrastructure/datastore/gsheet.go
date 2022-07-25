package datastore

type Store interface {
	Connect
}

type ServiceProvider struct {
	Token string
}
