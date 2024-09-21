package main

import (
	"backend/service/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	server      *http.Server
	router      *http.ServeMux
	userService user.UserService
}

func (s *Server) routes() {

	s.router.HandleFunc("/hel", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n"))
	})

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
		router: http.NewServeMux(),
	}

	server.routes()

	c := cors.New(cors.Options{})

	server.server.Handler = c.Handler(server.router)
	return &server

}
