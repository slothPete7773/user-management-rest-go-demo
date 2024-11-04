package handler

import (
	conduit "backend/model"
	"net/http"
)

type UserHandler struct {
	service conduit.UserService
}

func NewUserHandler(service conduit.UserService) UserHandler {
	return UserHandler{
		service,
	}
}

func (h UserHandler) MockUserPublicEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Public.\n"))
	}
}

func (h UserHandler) MockUserPrivateEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Private.\n"))
	}
}

func (h UserHandler) MockUserHealthcheckEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check.\n"))
	}
}
