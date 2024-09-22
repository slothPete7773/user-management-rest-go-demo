package userRepository

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	SayHi()
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{db}
}

func (repo userRepository) SayHi() {
	fmt.Println("Hello")
}
