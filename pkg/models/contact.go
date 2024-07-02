package models

type Contact struct {
	ID          uint   `gorm:"primaryKey,autoIncrement"`
	FullName    string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(255)"`
}

type ContactParam struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
