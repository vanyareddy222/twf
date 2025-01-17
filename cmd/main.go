package main

import (
	"log"
	"net/http"
	"twf1/internal/routes"
)

func CORSFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := routes.NewRouter()
	add := ":8080"
	log.Printf("Starting server at %v", "8080")
	// Apply CORS middleware
	err := http.ListenAndServe(add, CORSFilter(r))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
