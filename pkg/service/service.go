package service

import "github.com/Dponya/GreenDo/pkg/repository"

type Authorization interface {
}

type Greenist interface {
}

type GreenItem interface {
}

type Service struct {
	Authorization
	Greenist
	GreenItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
