package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/presenter"
	"github.com/rizalwildan/devcode-backend/pkg/models"
	"net/http"
)

func GetContacts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		contacts := []presenter.Contact{
			{ID: 3, FullName: "Supri", Email: "whassup@zyrex.com", PhoneNumber: "082211223344"},
			{ID: 5, FullName: "john", Email: "john@mail.com", PhoneNumber: "081023812083"},
		}

		return c.JSON(presenter.ContactListResponse(&contacts))
	}
}

func CreateContact() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody models.ContactParam
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}

		contact := &models.Contact{
			ID:          6,
			FullName:    requestBody.FullName,
			Email:       requestBody.Email,
			PhoneNumber: requestBody.PhoneNumber,
		}

		return c.JSON(presenter.ContactResponse(contact, "Contact created"))
	}
}
