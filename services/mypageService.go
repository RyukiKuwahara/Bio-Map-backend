package services

import (
	"fmt"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
)

func GetNameAndPosts(sessionId string) (string, []models.NewPost, error) {

	ur, err := repositories.NewUserRepository()
	if err != nil {
		return "", nil, err
	}
	userId, err := ur.GetUserId(sessionId)
	if err != nil {
		fmt.Println("GetUserId err")
		return "", nil, err
	}
	name, err := ur.GetName(userId)
	if err != nil {
		return "", nil, err
	}

	posts, err := ur.GetUserPosts(userId)
	if err != nil {
		return "", nil, err
	}
	fmt.Println(posts)

	newPosts, err := downloadImageFromFirebase(posts)
	if err != nil {
		return "", nil, err
	}

	return name, newPosts, err
}
