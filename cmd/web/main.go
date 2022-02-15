package main

import (
	"log"
	"net/http"

	"github.com/ansalamdaniel/library-api/pkg/models"
)

func main() {
	models.InitDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/books", getBooks)
	mux.HandleFunc("/books/create", createBook)

	log.Println("server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
