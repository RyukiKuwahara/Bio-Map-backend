package services

import (
	"fmt"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func createImagePath(userId, speciesId int, pr models.PostRequest) string {

	return fmt.Sprintf("user_id:%d_species_id:%d_lat:%f_lng:%f", userId, speciesId, pr.Lat, pr.Lng)
}

func Post(postRequest models.PostRequest) error {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		return err
	}
	userId, err := ur.GetUserId(postRequest.SessionId)
	if err != nil {
		return err
	}
	speciesId, err := ur.GetSpeciesId(postRequest.SpeciesName)
	if err != nil {
		return err
	}
	imagePath := createImagePath(userId, speciesId, postRequest)
	err = ur.RegisterPost(postRequest, userId, speciesId, imagePath)
	if err != nil {
		return err
	}
	return nil
}
