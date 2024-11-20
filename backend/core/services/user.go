package services

import (
	"backend/core/domain"
	"backend/core/ports"
	"fmt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return UserService{repo}
}

func (service UserService) NewUser(input domain.UserInput) (*domain.User, error) {
	user, err := service.repo.Create(domain.UserData{
		Name:           input.Name,
		FavoriteNumber: input.FavoriteNumber,
		HomeworldRealm: input.HomeworldRealm,
	})

	if err != nil {
		fmt.Println("service-new-user-error", err.Error())
	}
	return user, nil
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
