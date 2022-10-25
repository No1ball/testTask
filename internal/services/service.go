package services

import "github.com/No1ball/testEx/internal/repository"

type Followers interface {
}

type Templates interface {
}

type Api interface {
}
type Service struct {
	Followers
	Templates
	Api
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
