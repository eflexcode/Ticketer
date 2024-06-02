package database

import (
	"go.mod/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DbInstance *gorm.DB

func DBConnect() *gorm.DB {

	dbUrl := "host=localhost user=postgres password=Larry123 database=ticketer port=5432"

	dbInstance, error := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if error != nil {
		log.Fatal("Failed to connect to db")
	}
	dbInstance.AutoMigrate(&model.User{}, &model.Organisation{}, &model.Event{}, &model.Ticket{})
	return dbInstance

}
