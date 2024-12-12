package service

import (
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/repository"
)

type LimitService struct {
	repo *repository.LimitRepository
}

func NewLimitService(repo *repository.LimitRepository) *LimitService {
	return &LimitService{repo: repo}
}

func (s *LimitService) GetConsumerLimit(consumerID uint) ([]domain.Limit, error) {
	return s.repo.GetConsumerLimit(consumerID)
}

func (s *LimitService) CreateLimitsForConsumer(consumerID uint) ([]domain.Limit, error) {
	consumer, err := s.repo.GetConsumerByID(consumerID)
	if err != nil {
		return nil, err
	}

	// Hitung limit berdasarkan penghasilan
	limits := calculateLimitByIncome(consumer.Salary)

	// Kaitkan limit dengan konsumen
	for i := range limits {
		limits[i].ConsumerID = consumerID
	}

	err = s.repo.CreateLimits(limits)
	return limits, err
}

func calculateLimitByIncome(gaji float64) []domain.Limit {
	var limits []domain.Limit

	switch {
	case gaji < 3000000: // Penghasilan Rendah
		limits = []domain.Limit{
			{Tenor: 1, LimitAmount: 100000},
			{Tenor: 2, LimitAmount: 200000},
			{Tenor: 3, LimitAmount: 500000},
			{Tenor: 6, LimitAmount: 700000},
		}
	case gaji >= 3000000 && gaji < 7000000: // Penghasilan Menengah
		limits = []domain.Limit{
			{Tenor: 1, LimitAmount: 500000},
			{Tenor: 2, LimitAmount: 1000000},
			{Tenor: 3, LimitAmount: 1500000},
			{Tenor: 6, LimitAmount: 2000000},
		}
	case gaji >= 7000000 && gaji < 15000000: // Penghasilan Tinggi
		limits = []domain.Limit{
			{Tenor: 1, LimitAmount: 1000000},
			{Tenor: 2, LimitAmount: 2000000},
			{Tenor: 3, LimitAmount: 3000000},
			{Tenor: 6, LimitAmount: 4000000},
		}
	default: // Penghasilan Sangat Tinggi
		limits = []domain.Limit{
			{Tenor: 1, LimitAmount: 2000000},
			{Tenor: 2, LimitAmount: 4000000},
			{Tenor: 3, LimitAmount: 6000000},
			{Tenor: 6, LimitAmount: 8000000},
		}
	}

	return limits
}
