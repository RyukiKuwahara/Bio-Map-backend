package services

import (
	"log"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func createImagePath(pr models.PostRequest) (string, error) {

}

func Post(postRequest models.PostRequest) error {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	userId, err := ur.GetUserId(postRequest.SessionId)
	speciesId, err := ur.GetSpeciesId(postRequest.SpeciesName)
	imagePath, err := createImagePath(postRequest)
	err = ur.RegisterPost(postRequest, userId, speciesId, imagePath)
	if err != nil {
		return err
	}
	return nil
}
