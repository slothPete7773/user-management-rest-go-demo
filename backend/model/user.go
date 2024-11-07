package conduit

type UserRepository interface {
	Create(UserData) (User, error)
	Update(string, UserData) (User, error)
	DeleteById(string) error
	GetById(string) (User, error)
	GetMany() ([]User, error)
}

type UserService interface {
	NewUser(UserInput) (User, error)
	UpdateUser(string, UserInput) (User, error)
	DeleteUserById(string) error
	FindUserById(string) (User, error)
	FindMany() ([]User, error)
}

type User struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	FavoriteNumber int    `db:"favoriteNumber"`
	HomeworldRealm string `db:"homeworldRealm"`
}

type UserInput struct {
	Name           string `json:"name"`
	FavoriteNumber int    `json:"favoriteNumber"`
	HomeworldRealm string `db:"homeworldRealm"`
}

type UserData struct {
	Name           string `json:"name"`
	FavoriteNumber int    `json:"favoriteNumber"`
	HomeworldRealm string `db:"homeworldRealm"`
}

type UserResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	FavoriteNumber int    `json:"favoriteNumber"`
	HomeworldRealm string `db:"homeworldRealm"`
}
