package domain

import (
	"gorm.io/gorm"
)

type Limit struct {
	gorm.Model
	ConsumerID  uint     `gorm:"not null" json:"consumer_id"`
	Tenor       int      `json:"tenor"`
	LimitAmount float64  `json:"limit_amount"`
	Consumer    Consumer `gorm:"foreignKey:ConsumerID"`
}
