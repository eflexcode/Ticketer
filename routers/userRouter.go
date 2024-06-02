package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/model"
	"go.mod/util"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func getUser(id int, db *gorm.DB) (model.User, *gorm.DB) {
	var user model.User
	dbResult := db.Find(&user, "id = ?", id)
	return user, dbResult
}

func UserRouter(app *fiber.App, db *gorm.DB) {

	app.Post("/user/", func(ctx *fiber.Ctx) error {

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
		db.Create(&user)

		return ctx.Status(200).JSON(&user)
	})

	app.Get("/user/:id", func(ctx *fiber.Ctx) error {

		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.Status(400).JSON("Please insert valid id of user (int)")
		}

		user, dbResult := getUser(id, db)
		dbErr := dbResult.Error

		if dbErr != nil {
			return ctx.Status(500).JSON("Something went wrong")
		}

		if user.ID == 0 {
			errMessage := "No user found with id: " + strconv.Itoa(id)
			return ctx.Status(404).JSON(errMessage)
		}

		return ctx.Status(200).JSON(user)
	})

	app.Put("/user/:id", func(ctx *fiber.Ctx) error {

		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.Status(400).JSON("Please insert valid id of user (int)")
		}

		user, dbResult := getUser(id, db)
		dbErr := dbResult.Error

		if dbErr != nil {
			return ctx.Status(500).JSON("something went wrong")
		}

		if user.ID == 0 {
			return ctx.Status(400).JSON("No user found with id: " + strconv.Itoa(id))
		}

		var gottenUser util.User

		if err := ctx.BodyParser(&gottenUser); err != nil {
			return ctx.Status(400).JSON(err.Error())
		}

		db.Save(&user)

		return ctx.Status(200).JSON(user)

	})

}
