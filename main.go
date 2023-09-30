package main

import (
	"log"
	"net/http"

	"github.com/RyukiKuwahara/Bio-Map/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.CreateUserHandler)

	// cors := func(h http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		if r.Method == "OPTIONS" {
	// 			w.Header().Set("Access-Control-Allow-Origin", "*")
	// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 			return
	// 		}

	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 		h.ServeHTTP(w, r)
	// 	})
	// }

	// err := http.ListenAndServe(":8080", cors(http.DefaultServeMux))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
