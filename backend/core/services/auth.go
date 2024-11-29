package services

import (
	"backend/core/domain"
	"backend/core/ports"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	auth ports.AuthRepository
	user ports.UserRepository
}

func NewAuthService(auth ports.AuthRepository, user ports.UserRepository) ports.AuthService {
	return AuthService{auth, user}
}

func (svc AuthService) Login(input *domain.AuthCredential) (*domain.AuthResponse, error) {

	user, err := svc.user.GetUserById(input.Name)

	if err != nil {
		log.Println("error-login-find-user", err)
		return nil, err
	}

	log.Println("user password: ", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Println("error-login-compare-password-hash", err)
		return nil, errors.New("password-invalid")
	}

	token, err := svc.auth.SignUserAccessToken(&domain.AuthPassport{
		Id:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	})

	if err != nil {
		log.Println("error-login-generate-token", err)
		return nil, err
	}

	// response :=
	return &domain.AuthResponse{
		AccessToken: token,
	}, nil
}
