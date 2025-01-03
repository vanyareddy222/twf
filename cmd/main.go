package main

import (
	"log"
	"net/http"
	"twf1/internal/routes"
)

func main() {
	r := routes.NewRouter()
	add := ":8080"
	log.Printf("Starting server at %v", "8080")
	// Apply CORS middleware
	err := http.ListenAndServe(add, r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
