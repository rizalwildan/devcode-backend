package contact

import "github.com/rizalwildan/devcode-backend/pkg/models"

type Service interface {
	CreateContact(contact *models.ContactParam) (*models.Contact, error)
	GetContacts() ([]*models.Contact, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) CreateContact(contact *models.ContactParam) (*models.Contact, error) {
	data := &models.Contact{
		FullName:    contact.FullName,
		Email:       contact.Email,
		PhoneNumber: contact.PhoneNumber,
	}

	return s.repo.Create(data)
}

func (s service) GetContacts() ([]*models.Contact, error) {
	return s.repo.FetchAll()
}
