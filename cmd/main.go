package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/routers"
)

func main() {
	db := database.DBConnect()
	app := fiber.New()
	routers.InitDb(db)
	routers.UserRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	app.Listen(":3000")
}
