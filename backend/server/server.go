package server

import (
	"backend/adapter/handler"
	"backend/adapter/repository"
	"backend/core/domain"
	"backend/core/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	server      *http.Server
	router      *mux.Router
	userHandler *handler.UserHandler
	authhandler *handler.AuthHandler
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

	if gormDb, err := gorm.Open(sqlite.Open("database.db")); err == nil {
		gormDb.AutoMigrate(domain.User{})
	}

	db, err := sql.Open("sqlite3", "file:database.db?cache=shared")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Ping completed.")

	repository := repository.NewRepository(db)

	userService := services.NewUserService(repository)
	userHandler := handler.NewUserHandler(userService)

	// TODO: Refactor ASAP
	authService := services.NewAuthService(repository, repository)
	authHandler := handler.NewAuthHandler(authService)

	server.userHandler = userHandler
	server.authhandler = authHandler
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
