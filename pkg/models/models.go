package models

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	Isbn         string    `gorm:"primaryKey;index" json:"isbn"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	NoOfCopies   int       `json:"no_of_copies"`
	CallNumber   string    `json:"call_number"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func InitDb(db string) {
	var err error

	DB, err = gorm.Open(sqlite.Open(db), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB.AutoMigrate(&Book{})
}
