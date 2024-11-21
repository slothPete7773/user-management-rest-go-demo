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

func (svc UserService) NewUser(input domain.UserInput) (*domain.User, error) {
	user, err := svc.repo.Create(domain.UserData{
		Name:           input.Name,
		FavoriteNumber: input.FavoriteNumber,
		HomeworldRealm: input.HomeworldRealm,
	})

	if err != nil {
		fmt.Println("service-new-user-error", err.Error())
		return nil, err
	}
	return user, nil
}
func (svc UserService) UpdateUser(id string, input domain.UserInput) (*domain.User, error) {
	user, err := svc.repo.Update(id, domain.UserData{
		Name:           "slothPete",
		FavoriteNumber: 42,
		HomeworldRealm: "bakk",
	})

	if err != nil {
		fmt.Println("service-update-user-error", err.Error())
		return nil, err
	}
	return user, nil
}
func (UserService) DeleteUserById(string) error {
	return nil
}
func (UserService) FindUserById(string) (*domain.User, error) {
	return &domain.User{}, nil
}
func (svc UserService) FindMany() ([]*domain.User, error) {
	users, err := svc.repo.GetMany()
	if err != nil {
		fmt.Println("service-update-user-error", err.Error())
		return nil, err
	}
	return users, nil
}
