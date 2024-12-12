package domain

import (
	"gorm.io/gorm"
)

type Consumer struct {
	gorm.Model
	NIK          string        `gorm:"unique;not null" json:"nik"`
	FullName     string        `gorm:"not null" json:"full_name"`
	LegalName    string        `gorm:"not null" json:"legal_name"`
	BirthPlace   string        `json:"birth_place"`
	BirthDate    string        `json:"birth_date"`
	Salary       float64       `json:"salary"`
	KTPPhoto     string        `json:"ktp_photo"`
	SelfiePhoto  string        `json:"selfie_photo"`
	Transactions []Transaction `gorm:"foreignKey:ConsumerID"`
	Limits       []Limit       `gorm:"foreignKey:ConsumerID"`
}
