package domain

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	FavoriteNumber int    `json:"favorite_number"`
	HomeworldRealm string `json:"homeworld_realm"`
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
