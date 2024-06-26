package routers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/model"
	"go.mod/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbInstance *gorm.DB
var mongoClient *mongo.Client
var userCollection *mongo.Collection
var goCtx = context.TODO()

func InitDb(db *gorm.DB) {
	dbInstance = db
}
func InitMongoDb() {
	mongoClient = database.InitMongoDb()
	userCollection = database.GetCollection(mongoClient, "user")
}

//func getUser(id string) model.User {
//
//	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//
//	var user model.User
//
//	dbResult := dbInstance.Find(&user, "id = ?", id)
//	return user, dbResult
//}

func CreateUser(ctx *fiber.Ctx) error {

	var userGotten util.User

	if err := ctx.BodyParser(&userGotten); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var id = primitive.NewObjectID()

	timeNow := time.Time{}
	user := model.User{
		ID:        id,
		Email:     userGotten.Email,
		Username:  userGotten.Username,
		Password:  userGotten.Password,
		CreatedAt: timeNow,
	}

	_, err := userCollection.InsertOne(goCtx, &user)

	if err != nil {
		log.Fatal(err.Error())
		return ctx.Status(500).JSON("Something went wrong try again")
	}

	return ctx.Status(200).JSON(&user)
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of user")
	}

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(id)

	err := userCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(&user)
}
func GetUser2(id string) model.User {

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(id)

	_ = userCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&user)

	return user
}

func PutUser(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(id)

	err := userCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	var gottenUser util.User

	if err := ctx.BodyParser(&gottenUser); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if gottenUser.Email != "" {
		user.Email = gottenUser.Email
	}

	if gottenUser.Password != "" {
		user.Password = gottenUser.Password
	}

	if gottenUser.Username != "" {
		user.Username = gottenUser.Username
	}

	if gottenUser.ProfileImageUrl != "" {
		user.ProfileImageUrl = gottenUser.ProfileImageUrl
	}

	if gottenUser.CoverImageUrl != "" {
		user.CoverImageUrl = gottenUser.CoverImageUrl
	}

	update := bson.M{"email": user.Email, "password": user.Password, "username": user.Username, "profileimageurl": user.ProfileImageUrl, "coverimageUrl": user.CoverImageUrl}
	updateResult, err := userCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if updateResult.MatchedCount == 1 {
		return ctx.Status(200).JSON("Updated user id " + id)
	}

	return ctx.Status(500).JSON("something went wrong")
}
func PutUserTicket(id string, ticketId string) {

	var user model.User

	objId, _ := primitive.ObjectIDFromHex(id)

	_ = userCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&user)

	user.Tickets = append(user.Tickets, ticketId)
	update := bson.M{"tickets": user.Tickets}
	_, _ = userCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})
}
func BookmarkEvent(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	eventId := ctx.Params("event_id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of user")
	}

	if eventId == "" {
		return ctx.Status(400).JSON("Please insert valid id of event")
	}

	var user model.User
	var isEventIdValid = CheckIfEventIdIfValid(eventId)

	if !isEventIdValid {
		return ctx.Status(200).JSON("Invalid event id")
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	err := userCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	user.EventsInterestedIn = append(user.EventsInterestedIn, eventId)

	update := bson.M{"eventsinterestedin": user.EventsInterestedIn}
	updateResult, err := userCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if updateResult.MatchedCount == 1 {
		return ctx.Status(200).JSON("Event added to bookmarks")
	}

	return ctx.Status(500).JSON("something went wrong")
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of user (int)")
	}

	obj, _ := primitive.ObjectIDFromHex(id)

	result, err := userCollection.DeleteOne(goCtx, bson.M{"_id": obj})

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if result.DeletedCount < 1 {
		return ctx.Status(404).JSON("No user Found with id")
	}

	return ctx.Status(200).JSON("Deleted")
}

func UserRouter(app *fiber.App) {

	app.Post("/user/", CreateUser)
	app.Get("/user/:id", GetUser)
	app.Put("/user/:id", PutUser)
	app.Put("/user/bookmark/:id/:event_id", BookmarkEvent)
	app.Delete("/user/:id", DeleteUser)

}
