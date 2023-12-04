package main

import (
	"joglo-fiber-gorm/database"
	"joglo-fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	app := fiber.New()

	app.Get("/foo", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"foo": "bar",
		})
	})

	routes.RouterInit(app)

	app.Listen(":4545")
}
