package main

import (
	"backend/middlewares"
	"net/http"

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
		noAuthRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Health OK?\n"))
		})
		noAuthRouter.Handle("/public", s.userHandler.MockUserPublicEndpoint()).Methods("POST", "PUT", "GET")
	}

	mustAuthRouter := apiRouter.PathPrefix("").Subrouter()
	mustAuthRouter.Use(middlewares.Authenticate)
	{
		mustAuthRouter.Handle("/private", s.userHandler.MockUserPrivateEndpoint()).Methods("POST", "PUT", "GET")
		mustAuthRouter.Handle("/user", s.userHandler.MockUserPrivateEndpoint())
	}

}
