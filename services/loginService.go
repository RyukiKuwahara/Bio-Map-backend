package services

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func createSessionId() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	sessionID := hex.EncodeToString(randomBytes)

	return sessionID
}

// CreateUser creates a new user
func LoginUser(loginUser models.SigninUser) (string, error) {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	userId, err := ur.CheckUser(loginUser)
	if err != nil {
		return "", err
	}

	sessionId := createSessionId()
	err = ur.RegisterSessionId(sessionId, userId)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}
