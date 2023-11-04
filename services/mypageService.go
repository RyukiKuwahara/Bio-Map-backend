package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	firebase "firebase.google.com/go"
	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
	"google.golang.org/api/option"
)

func downloadBadgeFromFirebase(badgesPath []string) ([]models.Badge, error) {
	config := &firebase.Config{
		StorageBucket: "bio-map-storage.appspot.com",
	}
	opt := option.WithCredentialsFile("bio-map-storage-firebase-adminsdk-5lne1-e79313dcfa.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	var badgesData []models.Badge
	for _, badgePath := range badgesPath {
		rc, err := bucket.Object(badgePath).NewReader(ctx)
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		base64Img := base64.StdEncoding.EncodeToString(data)
		badgeData := models.Badge{ImageData: base64Img}
		badgesData = append(badgesData, badgeData)
	}
	return badgesData, nil
}

func GetUserInfo(sessionId string) (string, []models.NewPost, []models.Badge, error) {

	ur, err := repositories.NewUserRepository()
	if err != nil {
		return "", nil, nil, err
	}
	userId, err := ur.GetUserId(sessionId)
	if err != nil {
		fmt.Println("GetUserId err")
		return "", nil, nil, err
	}
	name, err := ur.GetName(userId)

	if err != nil {
		fmt.Println("GetName err")
		return "", nil, nil, err
	}

	posts, err := ur.GetUserPosts(userId)
	if err != nil {
		fmt.Println("GetUsePosts err")
		return "", nil, nil, err
	}

	badgesPath, err := ur.GetBadges(userId)
	if err != nil {
		fmt.Println("GetBadges err")
		return "", nil, nil, err
	}

	newPosts, err := downloadImageFromFirebase(posts)
	if err != nil {
		fmt.Println("downloadImage err")
		return "", nil, nil, err
	}

	badgesData, err := downloadBadgeFromFirebase(badgesPath)
	if err != nil {
		fmt.Println("downloadImage err")
		return "", nil, nil, err
	}

	return name, newPosts, badgesData, err
}
