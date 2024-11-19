package repository

import "backend/core/domain"

func (r Repository) Create(domain.UserData) (domain.User, error) {
	return domain.User{}, nil
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
