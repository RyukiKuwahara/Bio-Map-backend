package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
	"google.golang.org/api/option"
)

func createImagePath(userId, speciesId int, pr models.PostRequest) string {
	currentTime := time.Now()
	dateStr := currentTime.Format("2006-01-02")
	return fmt.Sprintf("posts/user_id:%d_species_id:%d_lat:%f_lng:%f_date:%s.jpg", userId, speciesId, pr.Lat, pr.Lng, dateStr)
}

func uploadImageToFirebase(base64Image, remoteFilename string) error {

	config := &firebase.Config{
		StorageBucket: "bio-map-storage.appspot.com",
	}

	opt := option.WithCredentialsFile("bio-map-storage-firebase-adminsdk-5lne1-e79313dcfa.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return err
	}

	contentType := "image/jpeg"

	decodedData, err := base64.StdEncoding.DecodeString(strings.Split(base64Image, ",")[1])
	if err != nil {
		return err
	}

	ctx := context.Background()
	writer := bucket.Object(remoteFilename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	_, err = writer.Write(decodedData)
	if err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	fmt.Println("Image uploaded to Firebase Storage.")
	return nil
}

func Post(postRequest models.PostRequest) error {
	ur, err := repositories.NewUserRepository()
	if err != nil {
		return err
	}
	userId, err := ur.GetUserId(postRequest.SessionId)
	if err != nil {
		fmt.Println("GetUserId err")
		return err
	}
	speciesId, err := ur.GetSpeciesId(postRequest.SpeciesName)
	if err != nil {
		fmt.Println("GetSpeciesId err")

		return err
	}
	imagePath := createImagePath(userId, speciesId, postRequest)
	err = ur.RegisterPost(postRequest, userId, speciesId, imagePath)
	if err != nil {
		fmt.Println("RegisterPost err")
		return err
	}

	err = uploadImageToFirebase(postRequest.ImageData, imagePath)
	if err != nil {
		return err
	}

	return nil
}
