package services

import "github.com/No1ball/testEx/internal/repository"

type TemplatesService struct {
	repos repository.TemplatesPostgres
}

func NewTemplatesService(repos repository.TemplatesPostgres) *TemplatesService {
	return &TemplatesService{repos: repos}
}
