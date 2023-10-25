package handlers

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/services"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var data models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newPosts, err := services.GetPosts(data.Name)
	if err != nil {
		http.Error(w, "Failed :  "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "[")
	for i, newPost := range newPosts {
		fmt.Fprintf(w, "{\"post_id\": %d, \"name\": \"%s\", \"image_data\":\"%x\", \"explain\":\"%s\", \"lat\":%f, \"lng\":%f}", newPost.PostId, newPost.SpeciesName, newPost.ImageData, newPost.Explain, newPost.Lat, newPost.Lng)
		if i < len(newPosts)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, "]")

}
