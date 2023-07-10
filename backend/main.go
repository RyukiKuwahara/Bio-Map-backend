package main

import (
	"log"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/backend/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.CreateUserHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
