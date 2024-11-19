package repository

import (
	"backend/core/ports"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.Repository {
	return Repository{db}
}
