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

	// addDummyData()
	log.Println("server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

// func addDummyData() {
// 	models.DB.Create(&models.Book{
// 		Isbn:         "978-0198520467",
// 		Title:        "Nuclear and Particle Physics",
// 		Author:       "W.S.C. Williams",
// 		NoOfCopies:   3,
// 		CallNumber:   "QC 776.W55.1991",
// 		Availability: "Available",
// 	})
// }
