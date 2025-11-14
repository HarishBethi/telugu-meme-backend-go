package main

import (
	"encoding/json"
	"net/http"
)

func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// Allow any origin for now (safe for public templates). If you later want to restrict, change this.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Allow common methods/headers for POST later
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setJSONHeader(w)
		response := map[string]string{"message": "Go backend working!"}
		json.NewEncoder(w).Encode(response)
	})

	// Templates API
	http.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		// handle OPTIONS preflight
		if r.Method == http.MethodOptions {
			setJSONHeader(w)
			w.WriteHeader(http.StatusOK)
			return
		}

		setJSONHeader(w)

		templates := []map[string]string{
			{
				"name": "Brahmi Krishna",
				"url":  "https://res.cloudinary.com/dui1jtybs/image/upload/v1763113242/brahmi-krishna_ptqsqu.gif",
			},
			// Add more templates here
		}

		json.NewEncoder(w).Encode(templates)
	})

	// Start server on port 8080
	http.ListenAndServe(":8080", nil)
}
