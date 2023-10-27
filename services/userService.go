package services

import (
	"log"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

// CreateUser creates a new user
func CreateUser(newUser models.SignupUser) error {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	err = ur.SaveUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
