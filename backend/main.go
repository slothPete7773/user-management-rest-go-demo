package main

import (
	"backend/server"
	"log"
)

func main() {
	srv := server.NewServer()
	log.Fatal(srv.Run(":8000"))

}
