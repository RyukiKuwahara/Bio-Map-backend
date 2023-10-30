package services

import (
	"context"
	"encoding/base64"
	"io/ioutil"

	firebase "firebase.google.com/go"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
	"google.golang.org/api/option"
)

func downloadImageFromFirebase(posts []models.Post) ([]models.NewPost, error) {
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

	// log.Printf("Downloaded contents: %v\n", string(data))

	ctx := context.Background()
	var newPosts []models.NewPost
	for _, post := range posts {
		rc, err := bucket.Object(post.ImagePath).NewReader(ctx)
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		base64Data := base64.StdEncoding.EncodeToString(data)
		newPost := models.NewPost{PostId: post.PostId, SpeciesName: post.SpeciesName, ImageData: base64Data, Explain: post.Explain, Lat: post.Lat, Lng: post.Lng}
		newPosts = append(newPosts, newPost)
	}
	return newPosts, nil
}

func GetPosts(name string) ([]models.NewPost, error) {

	ur, err := repositories.NewUserRepository()
	if err != nil {
		return nil, err
	}
	posts, err := ur.GetPosts(name)
	if err != nil {
		return nil, err
	}

	newPosts, err := downloadImageFromFirebase(posts)
	if err != nil {
		return nil, err
	}

	return newPosts, err
}

// func saveImage(data []byte, fileName string) error {
// 	err := ioutil.WriteFile(fileName, data, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
