package ports

import "backend/core/domain"

type UserRepository interface {
	Create(domain.UserData) (domain.User, error)
	Update(string, domain.UserData) (domain.User, error)
	DeleteById(string) error
	GetById(string) (domain.User, error)
	GetMany() ([]domain.User, error)
}

type UserService interface {
	NewUser(domain.UserInput) (domain.User, error)
	UpdateUser(string, domain.UserInput) (domain.User, error)
	DeleteUserById(string) error
	FindUserById(string) (domain.User, error)
	FindMany() ([]domain.User, error)
}
