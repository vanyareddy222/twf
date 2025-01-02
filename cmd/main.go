package main

import (
	"log"
	"net/http"
	"twf1/internal/routes"
)

func main() {
	http.HandleFunc("/calculate", routes.CalculateDeliveryCost)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
