package main

import (
	"backend/server"
	"fmt"
	"log"
)

var port = 8080

func main() {
	srv := server.NewServer()

	fmt.Println(fmt.Sprintf("Server running at http://localhost:%d", port))
	log.Fatal(srv.Run(fmt.Sprintf(":%d", port)))

}
