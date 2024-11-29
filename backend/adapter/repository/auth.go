package repository

import (
	"backend/core/domain"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (Repository) SignUserAccessToken(data *domain.AuthPassport) (string, error) {
	claims := domain.AuthClaims{
		Id:   data.Id,
		Name: data.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "user_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	signingKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Println("auth-error-sign-token", err.Error())
		return "", err
	}

	log.Println("ss: ", ss)
	return ss, nil
}
