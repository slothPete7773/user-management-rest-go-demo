package main

import (
	"backend/middlewares"
	"backend/service/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	server      *http.Server
	router      *mux.Router
	userService user.UserService
}

func (s *Server) routes() {

	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Adjust this based on your use case
	}).Handler)
	s.router.Use(middlewares.Logger)

	s.router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	noAuthRouter := apiRouter.PathPrefix("").Subrouter()
	{
		noAuthRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Health OK?\n"))
		})
		noAuthRouter.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
			id := r.URL.Query().Get("id")
			w.Write([]byte("id: " + id + "\n"))
		}).Methods("POST", "PUT", "GET")
	}

	mustAuthRouter := apiRouter.PathPrefix("").Subrouter()
	mustAuthRouter.Use(middlewares.Authenticate)
	{
		mustAuthRouter.HandleFunc("/private", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Private api\n"))
		}).Methods("GET")
	}

}

func (s *Server) run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	s.server.Addr = port
	log.Printf("server starting on %s", port)
	return s.server.ListenAndServe()
}

func main() {
	server := NewServer()
	err := server.run(":8000")
	if err != nil {
		panic(err.Error())
	}

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

	server.routes()
	server.server.Handler = server.router
	return &server

}
