package ports

import "backend/core/domain"

type AuthRepository interface {
	SignUserAccessToken(*domain.AuthPassport) (string, error)
}

type AuthService interface {
	Login(*domain.AuthCredential) (*domain.AuthResponse, error)
}
