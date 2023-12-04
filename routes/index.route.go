package routes

import (
	"joglo-fiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "OK DONK",
		})
	})
	r.Get("/users", controllers.GetUsers)
	r.Get("/users/:id", controllers.GetUserDetail)
	r.Post("/users", controllers.RegisterUser)
}
