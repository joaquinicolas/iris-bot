package entities

type Filter interface {
	Tag() string
}

type Product struct {
	Id    string
	Name  string
	Price string
}

func (p Product) Tag() string {
	return p.Name
}
