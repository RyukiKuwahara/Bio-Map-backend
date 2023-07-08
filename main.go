package main

import (
	"log"

	"github.com/RyukiKuwahara/Bio-Map/backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users", handlers.CreateUserHandler)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
