package models

import (
	"log"

	"github.com/Christomesh/gopostgres/db"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id,omitempty"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func SetUpBookMigration() *gorm.DB {
	db, err := db.NewConnection()
	if err != nil {
		log.Fatal("Failed to load database")
	}
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("could not migrate db")
	}
	return db
}
