package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/routes"
	"github.com/rizalwildan/devcode-backend/internal/mysql"
	"github.com/rizalwildan/devcode-backend/pkg/contact"
	"github.com/rizalwildan/devcode-backend/pkg/models"
)

func main() {
	app := fiber.New()

	db := mysql.Connect()
	err := db.AutoMigrate(&models.Contact{})
	if err != nil {
		return
	}

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello world"})
	})

	contactRepo := contact.NewRepository(db)
	contactService := contact.NewService(contactRepo)

	routes.ContactRouter(app, contactService)

	err = app.Listen(":3030")
	if err != nil {
		return
	}
}
