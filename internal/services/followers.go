package services

import "github.com/No1ball/testEx/internal/repository"

type FollowersService struct {
	repos repository.FollowersPostgres
}

func NewFollowerService(repos repository.FollowersPostgres) *FollowersService {
	return &FollowersService{repos: repos}
}
