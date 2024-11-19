package services

import (
	"backend/core/domain"
	"backend/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return UserService{repo}
}

func (UserService) NewUser(domain.UserInput) (domain.User, error) {
	return domain.User{}, nil
}
func (UserService) UpdateUser(string, domain.UserInput) (domain.User, error) {
	return domain.User{}, nil
}
func (UserService) DeleteUserById(string) error {
	return nil
}
func (UserService) FindUserById(string) (domain.User, error) {
	return domain.User{}, nil
}
func (UserService) FindMany() ([]domain.User, error) {
	return nil, nil
}
