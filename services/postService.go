package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/repositories"
	"github.com/nfnt/resize"
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

	decodedData, err := base64.StdEncoding.DecodeString(base64Image)
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

func resizeImage(base64Image string) (string, error) {
	parts := strings.Split(base64Image, ";base64,")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid data")
	}
	base64Image = parts[1]

	decoded, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return "", err
	}

	currentWidth := img.Bounds().Dx()
	currentHeight := img.Bounds().Dy()
	totalPixels := currentWidth * currentHeight
	maxTotalPixels := 90000 // 300 x 300 px程度
	percent := math.Sqrt(float64(maxTotalPixels) / float64(totalPixels))

	var newWidth, newHeight int

	if totalPixels <= maxTotalPixels {
		newWidth = currentWidth
		newHeight = currentHeight
	} else {
		newWidth = int(float64(currentWidth) * percent)
		newHeight = int(float64(currentHeight) * percent)
	}

	resizedImage := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, resizedImage, nil)
	if err != nil {
		return "", err
	}

	base64ResizedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64ResizedImage, nil
}

func NewError(s string) {
	panic("unimplemented")
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

	resizedImage, err := resizeImage(postRequest.ImageData)
	if err != nil {
		fmt.Println("resizeImage err")
		return err
	}

	err = uploadImageToFirebase(resizedImage, imagePath)
	if err != nil {
		return err
	}

	return nil
}
