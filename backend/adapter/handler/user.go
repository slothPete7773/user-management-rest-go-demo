package handler

import (
	"backend/adapter/utils"
	"backend/core/domain"
	"backend/core/ports"
	"fmt"
	"net/http"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service}
}

func (hl UserHandler) NewUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var u domain.UserInput
		err := utils.DecodeJSONBody(w, r, &u)
		if err != nil {
			fmt.Println("error-handler-request-body")
			switch val := err.(type) {
			case *utils.MalformedRequest:
				http.Error(w, val.Msg, val.Status)
			case error:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}

		fmt.Fprintf(w, "User: %v\n", u)
		// w.Write([]byte("NewUser\n"))
		hl.service.NewUser(domain.UserInput{
			Name:           "pete",
			FavoriteNumber: 773,
			HomeworldRealm: "Lopburi",
		})
	})
}
func (hl UserHandler) UpdateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateUser\n"))
		user, err := hl.service.UpdateUser("5b9f1668-c8b6-48e3-843a-1490e7f3e926", domain.UserInput{
			Name:           "slothPete",
			FavoriteNumber: 42,
			HomeworldRealm: "bakk",
		})

		if err != nil {
			fmt.Println("error-handler-update-user")
			http.Error(w, "Error", http.StatusInternalServerError)
		}

		fmt.Println("User: ", user)
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
func (hl UserHandler) FindMany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindMany\n"))
		users, err := hl.service.FindMany()

		if err != nil {
			fmt.Println("error-handler-update-user")
			http.Error(w, "Error", http.StatusInternalServerError)
		}

		for _, user := range users {
			fmt.Println("User: ", user)
		}

	})
}
