package contact

import (
	"github.com/rizalwildan/devcode-backend/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(contact *models.Contact) (*models.Contact, error)
	FetchAll() ([]*models.Contact, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(contact *models.Contact) (*models.Contact, error) {
	result := r.db.Create(&contact)
	if result.Error != nil {
		return nil, result.Error
	}

	return contact, nil
}

func (r repository) FetchAll() ([]*models.Contact, error) {
	var contacts []*models.Contact
	result := r.db.Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}

	return contacts, nil
}
