package services

import (
	"errors"
	"log"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func checkPassword(password string) error {
	if len(password) < 9 {
		return errors.New("パスワードの長さの長さが足りません．")
	}
	var hasLowL, hasUpL, hasNum bool
	for _, c := range password {
		if 'a' <= c && c <= 'z' {
			hasLowL = true
		} else if 'A' <= c && c <= 'Z' {
			hasUpL = true
		} else if '0' <= c && c <= '9' {
			hasNum = true
		}
	}
	if !(hasLowL && hasUpL && hasNum) {
		return errors.New("英小文字，英大文字，数字のいずれかが含まれていません．")
	}
	return nil
}

func CreateUser(newUser models.SignupUser) error {
	// Call the user repository to save the user in the database
	ur, err := repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}
	err = checkPassword(newUser.Password)
	if err != nil {
		return err
	}
	err = ur.SaveUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
