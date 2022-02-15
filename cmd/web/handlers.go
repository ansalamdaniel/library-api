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
			return
		}

		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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

func createBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var b models.Book

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := models.DB.Create(&b)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(struct {
			Message string `json:"message"`
		}{Message: "success"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(js)

	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
