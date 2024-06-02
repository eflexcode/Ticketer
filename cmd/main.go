package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/routers"
)

func main() {
	db := database.DBConnect()
	app := fiber.New()
	routers.UserRouter(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	app.Listen(":3000")
}
