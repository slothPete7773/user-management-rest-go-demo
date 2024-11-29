package server

import (
	"backend/adapter/middlewares"

	"github.com/rs/cors"
)

func (s *Server) routes() {

	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Adjust this based on your use case
	}).Handler)
	s.router.Use(middlewares.Logger)

	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()

	noAuthRouter := apiRouter.PathPrefix("").Subrouter()
	{
		noAuthRouter.Handle("/user/id", s.userHandler.FindUserById()).Methods("GET")
		noAuthRouter.Handle("/users", s.userHandler.FindMany()).Methods("GET")
	}

	mustAuthRouter := apiRouter.PathPrefix("").Subrouter()
	mustAuthRouter.Use(middlewares.Authenticate)
	{
		mustAuthRouter.Handle("/new-user", s.userHandler.NewUser()).Methods("POST")
		mustAuthRouter.Handle("/update-user", s.userHandler.UpdateUser()).Methods("PUT")
		mustAuthRouter.Handle("/delete-user", s.userHandler.DeleteUserById()).Methods("PUT")
		mustAuthRouter.Handle("/login", s.authhandler.Login()).Methods("POST")
	}

}
