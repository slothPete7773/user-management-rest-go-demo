package ports

import "backend/core/domain"

type UserRepository interface {
	CreateUser(domain.UserData) (*domain.User, error)
	UpdateUser(string, domain.UserData) (*domain.User, error)
	DeleteUserById(string) error
	GetUserById(string) (*domain.User, error)
	GetUserMany() ([]*domain.User, error)
}

type UserService interface {
	NewUser(domain.UserInput) (*domain.User, error)
	UpdateUser(string, domain.UserInput) (*domain.User, error)
	DeleteUserById(string) error
	FindUserById(string) (*domain.User, error)
	FindMany() ([]*domain.User, error)
}
