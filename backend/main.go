package main

import "log"

func main() {
	server := NewServer()
	log.Fatal(server.run(":8000"))

}
