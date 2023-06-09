package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"promotion/models"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Promotion{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Promotion{})
	if err != nil {
		return err
	}

	log.Println("Promotion table migrated")

	return nil
}
