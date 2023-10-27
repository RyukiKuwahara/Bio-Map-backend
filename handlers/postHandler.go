package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/services"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var postRequest models.PostRequest
	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = services.Post(postRequest)
	if err != nil {
		http.Error(w, "Failed to post. "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "post successfully")

}
