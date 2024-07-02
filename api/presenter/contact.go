package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalwildan/devcode-backend/pkg/models"
)

type Contact struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func ContactResponse(data *models.Contact, msg string) *fiber.Map {
	contact := Contact{
		ID:          data.ID,
		FullName:    data.FullName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
	}

	return &fiber.Map{
		"status":  "Success",
		"message": msg,
		"data":    contact,
	}
}

func ContactListResponse(data []*Contact) *fiber.Map {
	return &fiber.Map{
		"status": "Success",
		"data":   data,
	}
}

func ContactErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  "Failed",
		"message": err.Error(),
	}
}
