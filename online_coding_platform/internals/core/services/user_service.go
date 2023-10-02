package services

import (
	"github.com/online_coding_platform/internals/core/domains"
	"github.com/online_coding_platform/internals/core/ports"
)

type UserService struct {
	repo ports.UserRepo
}

func NewUserService(repo ports.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) CreateUser(name string) error {
	user := domains.NewUserWithDefScore(name)
	return service.repo.SaveUser(user)
}

func (service *UserService) AllUser() ([]domains.User, error) {
	return service.repo.GetAll()
}
