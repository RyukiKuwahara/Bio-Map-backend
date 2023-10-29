package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/handlers"
	"github.com/RyukiKuwahara/Bio-Map/setups"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Found!")
	fmt.Println("accuses /")
}

func main() {
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
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				return
			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			h.ServeHTTP(w, r)
		})
	}

	err := http.ListenAndServe(":8080", cors(http.DefaultServeMux))
	// err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	} else {
		fmt.Println("Server successed to start")
	}
}
