package models

import (
	"log"
	"os"
)

func SetupTestDb() {
	InitDb("test_db.sqlite3")
	tdbooks := []Book{
		{
			Isbn:         "978-0198520467",
			Title:        "Nuclear and Particle Physics",
			Author:       "W.S.C. Williams",
			NoOfCopies:   3,
			CallNumber:   "QC 776.W55.1991",
			Availability: "Available",
		},
		{
			Isbn:         "978-8124801604",
			Title:        "Emma",
			Author:       "Jane Austen",
			NoOfCopies:   8,
			CallNumber:   "823.7 AUS-E",
			Availability: "Available",
		},
	}
	for _, book := range tdbooks {
		DB.Create(&book)
	}
}

func TeardownTestDb() {
	sqlDb, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb.Close()

	// Cleaning up sqlite3 db file.
	err = os.Remove("test_db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
}
