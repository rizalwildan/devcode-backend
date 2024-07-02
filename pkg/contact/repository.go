package contact

import (
	"github.com/rizalwildan/devcode-backend/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(contact *models.Contact) (*models.Contact, error)
	FetchAll() ([]*models.Contact, error)
	Update(id uint, contact *models.Contact) (*models.Contact, error)
	Delete(id uint) error
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

func (r repository) Update(id uint, contact *models.Contact) (*models.Contact, error) {
	var contactModel models.Contact
	result := r.db.Model(&contactModel).Clauses(clause.Returning{}).Where("id = ?", id).UpdateColumns(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	contactModel.ID = id

	return &contactModel, nil
}

func (r repository) Delete(id uint) error {
	if errModel := r.db.Delete(&models.Contact{}, id); errModel.Error != nil {
		return errModel.Error
	}

	return nil
}
