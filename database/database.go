package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBConnect() {

	dbUrl := "host=localhost user=postgres password=Larry123 dbName=ticketer port=5432"

	db, error := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if error != nil {
		log.Fatal("Failed to connect to db")
	}

	db.AutoMigrate()
}
