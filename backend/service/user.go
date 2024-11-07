package service

import conduit "backend/model"

type userService struct {
	repository conduit.UserRepository
}

func NewUserService(repository conduit.UserRepository) conduit.UserService {
	return userService{
		repository,
	}
}

func (s userService) NewUser(input conduit.UserInput) (conduit.User, error) {
	return conduit.User{}, nil
}
func (s userService) UpdateUser(id string, input conduit.UserInput) (conduit.User, error) {
	return conduit.User{}, nil
}
func (s userService) DeleteUserById(id string) error {
	return nil
}
func (s userService) FindUserById(id string) (conduit.User, error) {
	return conduit.User{}, nil
}
func (s userService) FindMany() ([]conduit.User, error) {
	return nil, nil
}
