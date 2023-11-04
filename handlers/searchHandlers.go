package handlers

import (
	"encoding/json"

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
	requestJson, err := json.Marshal(newPosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(requestJson)
}
