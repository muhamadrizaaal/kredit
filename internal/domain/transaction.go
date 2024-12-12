package domain

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ConsumerID     uint     `gorm:"not null" json:"consumer_id"`
	ContractNumber string   `gorm:"unique;not null" json:"contract_number"`
	OTR            float64  `json:"otr"`
	AdminFee       float64  `json:"admin_fee"`
	Installment    float64  `json:"installment"`
	Interest       float64  `json:"interest"`
	AssetName      string   `json:"asset_name"`
	Consumer       Consumer `gorm:"foreignKey:ConsumerID"`
}
