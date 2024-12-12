package tests

import (
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockTransactionRepository adalah mock untuk repository transaksi
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *domain.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	// Buat mock repository
	mockTransactionRepo := new(MockTransactionRepository)
	mockConsumerRepo := new(MockConsumerRepository)

	// Buat service dengan mock repository
	transactionService := service.NewTransactionService(mockTransactionRepo, mockConsumerRepo)

	// Siapkan data konsumen dan transaksi untuk test
	consumer := &domain.Consumer{
		Model: gorm.Model{ID: 1},
		NIK:   "1234567890",
		Limits: []domain.Limit{
			{
				Tenor:       3,
				LimitAmount: 1000000,
			},
		},
	}

	transaction := &domain.Transaction{
		ConsumerID:     1,
		ContractNumber: "CONTRACT001",
		OTR:            500000,
		Installment:    3,
	}

	// Set expectation
	mockConsumerRepo.On("FindByID", uint(1)).Return(consumer, nil)
	mockTransactionRepo.On("Create", transaction).Return(nil)

	// Jalankan test
	err := transactionService.CreateTransaction(transaction)

	// Assertion
	assert.NoError(t, err)
	mockConsumerRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}
