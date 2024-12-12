package repository

import (
	"pt-xyz-multifinance/internal/domain"

	"gorm.io/gorm"
)

type ConsumerRepository struct {
	db *gorm.DB
}

func NewConsumerRepository(db *gorm.DB) *ConsumerRepository {
	return &ConsumerRepository{db: db}
}

func (r *ConsumerRepository) Create(consumer *domain.Consumer) error {
	return r.db.Create(consumer).Error
}

func (r *ConsumerRepository) FindByID(id uint) (*domain.Consumer, error) {
	var consumer domain.Consumer
	err := r.db.First(&consumer, id).Error
	return &consumer, err
}
