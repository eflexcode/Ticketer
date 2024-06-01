package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
)

func main() {
	database.DBConnect()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	app.Listen(":3000")
}
