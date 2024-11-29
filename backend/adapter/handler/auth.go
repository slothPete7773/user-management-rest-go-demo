package handler

import (
	"backend/adapter/utils"
	"backend/core/domain"
	"backend/core/ports"
	"log"
	"net/http"
)

type AuthHandler struct {
	service ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (hl AuthHandler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var a domain.AuthCredential
		err := utils.DecodeJSONBody(w, r, &a)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}

		response, err := hl.service.Login(&a)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}

		log.Println("login response: ", response)
	})
}
