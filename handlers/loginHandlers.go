package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/models"
	"github.com/RyukiKuwahara/Bio-Map/services"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser models.SigninUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session_id, err := services.LoginUser(loginUser)
	if err != nil {
		http.Error(w, "Failed to login user. "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"session_id\": \"%s\", \"message\": \"Success to login user\"}", session_id)

}
