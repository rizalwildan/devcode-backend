package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/handlers"
	"github.com/rizalwildan/devcode-backend/pkg/contact"
)

func ContactRouter(app fiber.Router, service contact.Service) {
	app.Get("/contacts", handlers.GetContacts(service))
	app.Post("/contacts", handlers.CreateContact(service))
}
