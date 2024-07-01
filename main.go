package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello world"})
	})

	err := app.Listen(":3030")
	if err != nil {
		return
	}
}
