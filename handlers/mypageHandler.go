package handlers

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/services"
)

func MypageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mypage")
	var data models.MypageRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	name, posts, badgesData, err := services.GetUserInfo(data.SessionId)
	if err != nil {
		http.Error(w, "Failed :  "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"name\": \"%s\", \"posts\": [", name)
	for i, post := range posts {
		fmt.Fprintf(w, "{\"post_id\": %d, \"name\": \"%s\", \"image_data\":\"%s\", \"explain\":\"%s\", \"lat\":%f, \"lng\":%f}", post.PostId, post.SpeciesName, post.ImageData, post.Explain, post.Lat, post.Lng)
		if i < len(posts)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, "], \"badges\": [")
	for i, badgeData := range badgesData {
		fmt.Println(badgeData)
		fmt.Fprintf(w, "{\"image_data\": \"%s\"}", badgeData)
		if i < len(badgesData)-1 {
			fmt.Fprintf(w, ", ")
		}
	}

	fmt.Fprintf(w, "]}")

}
