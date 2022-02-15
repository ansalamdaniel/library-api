package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ansalamdaniel/library-api/pkg/models"
)

func ping(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(struct {
		Message string
	}{Message: "pong!!!!!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("GET method on book api")
		var books []models.Book

		result := models.DB.Find(&books)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}

		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
