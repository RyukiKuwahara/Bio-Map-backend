package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/services"
)

func LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	var sessionId string
	err := json.NewDecoder(r.Body).Decode(&sessionId)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = services.LogoutUser(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"message\": \"Success to logout user\"}")

}
