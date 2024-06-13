package routers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/model"
	"go.mod/util"
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

func getUser(id int) (model.User, *gorm.DB) {
	var user model.User
	dbResult := dbInstance.Find(&user, "id = ?", id)
	return user, dbResult
}

func CreateUser(ctx *fiber.Ctx) error {

	var userGotten util.User

	if err := ctx.BodyParser(&userGotten); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	timeNow := time.Time{}
	user := model.User{
		Email:     userGotten.Email,
		Username:  userGotten.Username,
		Password:  userGotten.Password,
		CreatedAt: timeNow,
	}

	//dbInstance.Create(&user)

	result, err := userCollection.InsertOne(goCtx, &user)

	if err != nil {
		log.Fatal(err.Error())
		return ctx.Status(500).JSON("Something went wrong try again")
	}

	//user.Id = result.InsertedID
	//
	//var s string = result.InsertedID

	//	strconv.presult)

	return ctx.Status(200).JSON(result)
}

//func GetUser(ctx *fiber.Ctx) error {
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of user (int)")
//	}
//
//	user, dbResult := getUser(id)
//	dbErr := dbResult.Error
//
//	if dbErr != nil {
//		return ctx.Status(500).JSON("Something went wrong")
//	}
//
//	if user.ID == 0 {
//		errMessage := "No user found with id: " + strconv.Itoa(id)
//		return ctx.Status(404).JSON(errMessage)
//	}
//
//	return ctx.Status(200).JSON(&user)
//}
//
//func PutUser(ctx *fiber.Ctx) error {
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of user (int)")
//	}
//
//	dbUser, dbResult := getUser(id)
//	dbErr := dbResult.Error
//
//	if dbErr != nil {
//		return ctx.Status(500).JSON("Something went wrong")
//	}
//
//	if dbUser.ID == 0 {
//		return ctx.Status(400).JSON("No user found with id: " + strconv.Itoa(id))
//	}
//
//	var gottenUser util.User
//
//	if err := ctx.BodyParser(&gottenUser); err != nil {
//		return ctx.Status(500).JSON(err.Error())
//	}
//
//	if gottenUser.Email != "" {
//		dbUser.Email = gottenUser.Email
//	}
//
//	if gottenUser.Password != "" {
//		dbUser.Password = gottenUser.Password
//	}
//
//	if gottenUser.Username != "" {
//		dbUser.Username = gottenUser.Username
//	}
//
//	if gottenUser.ProfileImageUrl != "" {
//		dbUser.ProfileImageUrl = gottenUser.ProfileImageUrl
//	}
//
//	if gottenUser.CoverImageUrl != "" {
//		dbUser.CoverImageUrl = gottenUser.CoverImageUrl
//	}
//
//	dbInstance.Save(&dbUser)
//
//	return ctx.Status(200).JSON(&dbUser)
//}
//
//func DeleteUser(ctx *fiber.Ctx) error {
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of user (int)")
//	}
//
//	user, db := getUser(id)
//	dbError := db.Error
//
//	if dbError != nil {
//		return ctx.Status(500).JSON("Something went wrong")
//	}
//
//	if user.ID == 0 {
//		return ctx.Status(400).JSON("No user found with id: " + strconv.Itoa(id))
//	}
//
//	db.Delete(&user)
//
//	return ctx.Status(200).JSON("Deleted")
//}

func UserRouter(app *fiber.App) {

	app.Post("/user/", CreateUser)
	//app.Get("/user/:id", GetUser)
	//app.Put("/user/:id", PutUser)
	//app.Delete("/user/:id", DeleteUser)

}
