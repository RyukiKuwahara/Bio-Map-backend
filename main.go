package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RyukiKuwahara/Bio-Map/handlers"
	"github.com/RyukiKuwahara/Bio-Map/setups"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Found!")
	fmt.Println("accuses /")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	fmt.Println("starting main")

	setups.Initialization()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", handlers.CreateUserHandler)
	http.HandleFunc("/login", handlers.LoginUserHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/mypage", handlers.MypageHandler)

	cors := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				return
			}

			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			h.ServeHTTP(w, r)
		})
	}

	err = http.ListenAndServe(":8080", cors(http.DefaultServeMux))
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	} else {
		fmt.Println("Server successed to start")
	}
}
