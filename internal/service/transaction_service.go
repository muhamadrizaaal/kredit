package service

import (
	"errors"
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/repository"
	"sync"
)

type TransactionService struct {
	repo         *repository.TransactionRepository
	consumerRepo *repository.ConsumerRepository
	mu           sync.Mutex
}

func NewTransactionService(
	repo *repository.TransactionRepository,
	consumerRepo *repository.ConsumerRepository,
) *TransactionService {
	return &TransactionService{
		repo:         repo,
		consumerRepo: consumerRepo,
	}
}

func (s *TransactionService) CreateTransaction(transaction *domain.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validasi konsumen
	consumer, err := s.consumerRepo.FindByID(transaction.ConsumerID)
	if err != nil {
		return errors.New("konsumen tidak ditemukan")
	}

	// Validasi limit
	var totalTransactionAmount float64
	for _, t := range consumer.Transactions {
		totalTransactionAmount += t.OTR
	}

	// Cek limit
	var consumerLimit *domain.Limit
	for _, l := range consumer.Limits {
		if float64(l.Tenor) == transaction.Installment {
			consumerLimit = &l
			break
		}
	}

	if consumerLimit == nil {
		return errors.New("limit tidak ditemukan")
	}

	if totalTransactionAmount+transaction.OTR > consumerLimit.LimitAmount {
		return errors.New("melebihi limit transaksi")
	}

	// Buat transaksi
	return s.repo.Create(transaction)
}

func (s *TransactionService) GetTransaction(id uint) (*domain.Transaction, error) {
	return s.repo.FindByID(id)
}
