package services

import (
	"backend/core/domain"
	"backend/core/ports"
	"fmt"
)

type UserService struct {
	user ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return UserService{repo}
}

func (svc UserService) NewUser(input domain.UserInput) (*domain.User, error) {
	user, err := svc.user.CreateUser(domain.UserData{
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
	user, err := svc.user.UpdateUser(id, domain.UserData{
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
func (svc UserService) DeleteUserById(id string) error {
	err := svc.user.DeleteUserById(id)
	if err != nil {
		fmt.Println("service-delete-user-error", err.Error())
		return err
	}
	return nil
}
func (svc UserService) FindUserById(id string) (*domain.User, error) {

	user, err := svc.user.GetUserById(id)
	if err != nil {
		fmt.Println("service-get-user-by-id-error", err.Error())
		return nil, err
	}
	return user, nil
}
func (svc UserService) FindMany() ([]*domain.User, error) {
	users, err := svc.user.GetUserMany()
	if err != nil {
		fmt.Println("service-update-user-error", err.Error())
		return nil, err
	}
	return users, nil
}
