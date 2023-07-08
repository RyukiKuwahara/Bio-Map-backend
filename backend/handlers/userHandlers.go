package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/backend/models"
	"github.com/RyukiKuwahara/Bio-Map/backend/services"
)

// CreateUserHandler handles the user registration endpoint
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call user service to create user
	err = services.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}
