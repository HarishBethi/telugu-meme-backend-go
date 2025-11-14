package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Go backend working!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Templates API
	http.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {

		templates := []map[string]string{
			{
				"name": "Brahmi Krishna",
				"url":  "https://res.cloudinary.com/dui1jtybs/image/upload/v1763113242/brahmi-krishna_ptqsqu.gif",
			},
			// Add more templates here as you upload them to Cloudinary
			// {
			//   "name": "Another Template",
			//   "url":  "https://res.cloudinary.com/dui1jtybs/image/upload/.../template.jpg",
			// },
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(templates)
	})

	// Start server on port 8080
	http.ListenAndServe(":8080", nil)
}
