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
	var data models.SessionData
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

	response := models.MypageResponse{
		Name:       name,
		Posts:      posts,
		BadgesData: badgesData,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
