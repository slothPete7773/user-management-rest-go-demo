package server

import (
	"backend/adapter/handler"
	"backend/adapter/repository"
	"backend/core/services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	server      *http.Server
	router      *mux.Router
	userHandler *handler.UserHandler
}

func NewServer() *Server {
	server := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
		router: mux.NewRouter().StrictSlash(true),
	}

	repository := repository.NewRepository(nil)

	userService := services.NewUserService(repository)
	userHandler := handler.NewUserHandler(userService)

	server.userHandler = userHandler
	server.routes()
	server.server.Handler = server.router
	return &server

}

func (s *Server) Run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	s.server.Addr = port
	log.Printf("server starting on %s", port)
	return s.server.ListenAndServe()
}
