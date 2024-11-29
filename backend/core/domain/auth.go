package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthCredential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthPassport struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

type AuthClaims struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	jwt.RegisteredClaims
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
