package conduit

type UserRepository interface {
	Create() error
}

type UserService interface {
	NewUser() error
}
