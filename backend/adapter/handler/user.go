package handler

import (
	"backend/core/domain"
	"backend/core/ports"
	"net/http"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service}
}

func (hlr UserHandler) NewUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NewUser\n"))
		hlr.service.NewUser(domain.UserInput{
			Name:           "pete",
			FavoriteNumber: 773,
			HomeworldRealm: "Lopburi",
		})
	})
}
func (UserHandler) UpdateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateUser\n"))
	})
}
func (UserHandler) DeleteUserById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeleteUserById\n"))
	})
}
func (UserHandler) FindUserById() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindUserById\n"))
	})
}
func (UserHandler) FindMany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindMany\n"))
	})
}
