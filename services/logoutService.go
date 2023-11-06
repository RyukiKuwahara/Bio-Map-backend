package services

import (
	"log"

	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func LogoutUser(sessionId string) error {
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	err = ur.RemoveSessionId(sessionId)
	if err != nil {
		return err
	}
	return nil
}
