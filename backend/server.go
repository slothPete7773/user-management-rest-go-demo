package main

import (
	"backend/handler"
	"backend/repository/postres"
	"backend/service"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	server      *http.Server
	router      *mux.Router
	userHandler handler.UserHandler
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

	userRepo := postres.NewUserRepository(nil)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	server.userHandler = userHandler
	server.routes()
	server.server.Handler = server.router
	return &server

}

func (s *Server) run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	s.server.Addr = port
	log.Printf("server starting on %s", port)
	return s.server.ListenAndServe()
}
