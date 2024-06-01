package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/model"
	"go.mod/util"
	"time"
)

func UserRouter(app *fiber.App) {

	app.Post("/user/", func(ctx *fiber.Ctx) error {

		var userGotten util.User

		if err := ctx.BodyParser(&userGotten); err != nil {
			return ctx.Status(400).JSON(err.Error())
		}

		timeNow := time.Time{}
		user := model.User{
			Email:     userGotten.Email,
			Username:  userGotten.Username,
			CreatedAt: timeNow,
		}
		database.DbInstance.Create(user)

		return ctx.Status(200).JSON(&user)
	})

}
