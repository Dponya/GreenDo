package repository

type Authorization interface {
}

type Greenist interface {
}

type GreenItem interface {
}

type Repository struct {
	Authorization
	Greenist
	GreenItem
}

func NewRepository() *Repository {
	return &Repository{}
}
