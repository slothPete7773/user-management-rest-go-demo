package domain

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	FavoriteNumber int       `json:"favorite_number"`
	HomeworldRealm string    `json:"homeworld_realm"`
}

type UserInput struct {
	Name           string `json:"name"`
	FavoriteNumber int    `json:"favorite_number"`
	HomeworldRealm string `json:"homeworld_realm"`
}

type UserData struct {
	Name           string
	FavoriteNumber int
	HomeworldRealm string
}
