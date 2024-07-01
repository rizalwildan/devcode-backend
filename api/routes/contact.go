package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/handlers"
)

func ContactRouter(app fiber.Router) {
	app.Get("/contacts", handlers.GetContacts())
	app.Post("/contacts", handlers.CreateContact())
}
