package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/routes"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello world"})
	})

	routes.ContactRouter(app)

	err := app.Listen(":3030")
	if err != nil {
		return
	}
}
