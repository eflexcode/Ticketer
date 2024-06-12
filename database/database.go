package database

import (
	"context"
	"go.mod/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DbInstance *gorm.DB

var cxt = context.TODO()

func DBConnect() *gorm.DB {

	dbUrl := "host=localhost user=postgres password=Larry123 database=ticketer port=5432"

	dbInstance, error := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if error != nil {
		log.Fatal("Failed to connect to db")
	}
	dbInstance.AutoMigrate(&model.User{}, &model.Organisation{}, &model.Event{}, &model.Ticket{})
	return dbInstance

}

func InitMongoDb() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongo://localhost:27017")
	client, err := mongo.Connect(cxt, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	var errPing = client.Ping(cxt, nil)

	if errPing != nil {
		log.Fatal(errPing)
	}

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Ticketer").Collection(collectionName)
	return collection
}
