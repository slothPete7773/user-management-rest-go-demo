package repository

import (
	"backend/core/domain"
	"fmt"

	"github.com/google/uuid"
)

func (r Repository) Create(data domain.UserData) (*domain.User, error) {
	query := `
	insert into users ( id, name, favorite_number, homeworld_realm)
	values (?, ?, ?, ?)
	`

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("error-user-create-tx", err.Error())
		return nil, err
	}

	id := uuid.New().String()
	_, err = tx.Exec(query, id, data.Name, data.FavoriteNumber, data.HomeworldRealm)
	if err != nil {
		fmt.Println("error-user-create", err.Error())
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error-user-create-commit", err.Error())
		return nil, err
	}
	return &domain.User{
		ID:             id,
		Name:           data.Name,
		FavoriteNumber: data.FavoriteNumber,
		HomeworldRealm: data.HomeworldRealm,
	}, nil
}
func (r Repository) Update(string, domain.UserData) (domain.User, error) {
	return domain.User{}, nil
}
func (r Repository) DeleteById(string) error {
	return nil
}
func (r Repository) GetById(string) (domain.User, error) {
	return domain.User{}, nil
}
func (r Repository) GetMany() ([]domain.User, error) {
	return nil, nil
}
