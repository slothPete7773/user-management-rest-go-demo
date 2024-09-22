package userHandler

import (
	userService "backend/modules/user/service"
	"log"
	"net/http"
)

// type UserHandler interface{}

type UserHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) UserHandler {
	return UserHandler{userService}
}

func (service UserHandler) Greet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service.userService.Greet()
		log.Println("Done greet() service.")

	}
}
