package postres

import (
	conduit "backend/model"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) conduit.UserRepository {
	return userRepository{
		db: db,
	}
}

func (r userRepository) Create() error {
	return nil
}
