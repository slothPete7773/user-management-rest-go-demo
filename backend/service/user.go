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

func (s userService) NewUser() error {
	s.repository.Create()
	return nil
}
