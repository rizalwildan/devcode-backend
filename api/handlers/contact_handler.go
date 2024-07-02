package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/api/presenter"
	"github.com/rizalwildan/devcode-backend/pkg/contact"
	"github.com/rizalwildan/devcode-backend/pkg/models"
	"net/http"
)

func GetContacts(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contacts, err := service.GetContacts()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}

		var result []*presenter.Contact
		for _, val := range contacts {
			result = append(result, &presenter.Contact{
				ID:          val.ID,
				FullName:    val.FullName,
				Email:       val.Email,
				PhoneNumber: val.PhoneNumber,
			})
		}

		return c.JSON(presenter.ContactListResponse(result))
	}
}

func CreateContact(service contact.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody models.ContactParam
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ContactErrorResponse(err))
		}

		result, err := service.CreateContact(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ContactErrorResponse(err))
		}

		return c.JSON(presenter.ContactResponse(result, "Contact created"))
	}
}
