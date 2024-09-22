package userService

import userRepository "backend/modules/user/repository"

type UserService interface {
	Greet()
}

type userService struct {
	userRepository userRepository.UserRepository
}

func NewUserService(userRepository userRepository.UserRepository) UserService {
	return userService{userRepository: userRepository}
}

func (service userService) Greet() {
	service.userRepository.SayHi()
}
