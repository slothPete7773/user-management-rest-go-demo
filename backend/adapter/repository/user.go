package repository

import (
	"backend/core/domain"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (r Repository) CreateUser(data domain.UserData) (*domain.User, error) {
	query := `
	insert into users ( id, password, name, favorite_number, homeworld_realm)
	values (?, ?, ?, ?, ?)
	`

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("error-user-create-tx", err.Error())
		return nil, err
	}

	id := uuid.New()
	passwordMock := uuid.NewString()
	_, err = tx.Exec(query, id.String(), passwordMock, data.Name, data.FavoriteNumber, data.HomeworldRealm)
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
		Password:       passwordMock,
		Name:           data.Name,
		FavoriteNumber: data.FavoriteNumber,
		HomeworldRealm: data.HomeworldRealm,
	}, nil
}
func (r Repository) UpdateUser(id string, data domain.UserData) (*domain.User, error) {
	query := `
	update users
	set
		name = ?,
		favorite_number = ?,
		homeworld_realm = ?
	where id = ?
	`

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("error-user-update-tx", err.Error())
		return nil, err
	}

	_, err = tx.Exec(query, data.Name, data.FavoriteNumber, data.HomeworldRealm, id)
	if err != nil {
		fmt.Println("error-user-update", err.Error())
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error-user-update-commit", err.Error())
		return nil, err
	}
	return &domain.User{
		ID:             uuid.MustParse(id),
		Name:           data.Name,
		FavoriteNumber: data.FavoriteNumber,
		HomeworldRealm: data.HomeworldRealm,
	}, nil
	// return domain.User{}, nil
}
func (r Repository) DeleteUserById(id string) error {
	query := `
	DELETE FROM customers WHERE customer_id = ?
	`

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("error-user-delete-tx", err.Error())
		return err
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		fmt.Println("error-user-delete", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error-user-delete-commit", err.Error())
		return err
	}
	return nil
}
func (r Repository) GetUserById(id string) (*domain.User, error) {
	query := `
	select * from users
	where name = ?
	`

	row := r.db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		fmt.Println("error-user-query-user-by-id", err)
		return nil, err
	}

	// var user domain.User
	user := new(domain.User)
	err := row.Scan(&user.ID, &user.Name, &user.FavoriteNumber, &user.HomeworldRealm)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	} else if err != nil {
		fmt.Println("error-user-query-user-by-id", err)
		return nil, err
	}

	return user, nil
}
func (r Repository) GetUserMany() ([]*domain.User, error) {
	query := `
	select * from users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("error-user-query-many-users", err.Error())
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		user := new(domain.User)
		err := rows.Scan(&user.ID, &user.Name, &user.FavoriteNumber, &user.HomeworldRealm)
		if err != nil {
			fmt.Println("error-user-query-many-users", err.Error())
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
