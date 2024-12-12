package service

import (
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/repository"
)

type ConsumerService struct {
	repo *repository.ConsumerRepository
}

func NewConsumerService(repo *repository.ConsumerRepository) *ConsumerService {
	return &ConsumerService{repo: repo}
}

func (s *ConsumerService) CreateConsumer(consumer *domain.Consumer) error {
	return s.repo.Create(consumer)
}

func (s *ConsumerService) GetConsumer(id uint) (*domain.Consumer, error) {
	return s.repo.FindByID(id)
}
