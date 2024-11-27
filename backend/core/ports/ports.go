package ports

type Repository interface {
	UserRepository
	AuthRepository
}
