package services

import (
	"github.com/RyukiKuwahara/Bio-Map/backend/models"
	"github.com/RyukiKuwahara/Bio-Map/backend/repositories"
)

// CreateUser creates a new user
func CreateUser(newUser models.User) error {
	// Call the user repository to save the user in the database
	err := repositories.SaveUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
