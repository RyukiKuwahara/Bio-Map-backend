package services

import (
	"fmt"
	"log"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

// CreateUser creates a new user
func LoginUser(loginUser models.SigninUser) error {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	err = ur.CheckUser(loginUser)
	if err != nil {
		return err
	}
	return nil
}
