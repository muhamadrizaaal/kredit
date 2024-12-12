package repository

import (
	"pt-xyz-multifinance/internal/domain"

	"gorm.io/gorm"
)

type LimitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) *LimitRepository {
	return &LimitRepository{db: db}
}

func (r *LimitRepository) GetConsumerLimit(consumerID uint) ([]domain.Limit, error) {
	var limits []domain.Limit
	err := r.db.Preload("Consumer").Where("consumer_id = ?", consumerID).Find(&limits).Error
	return limits, err
}

func (r *LimitRepository) CreateLimits(limits []domain.Limit) error {
	return r.db.Create(&limits).Error
}

func (r *LimitRepository) GetConsumerByID(consumerID uint) (*domain.Consumer, error) {
	var consumer domain.Consumer
	result := r.db.Find(&consumer, consumerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &consumer, nil
}
