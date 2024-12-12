package tests

import (
	"pt-xyz-multifinance/internal/domain"
	"pt-xyz-multifinance/internal/repository"
	"pt-xyz-multifinance/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Pastikan implementasi interface
var _ repository.ConsumerRepository = &MockConsumerRepository{}

type MockConsumerRepository struct {
	mock.Mock
}

// Implementasi method-method dari ConsumerRepository interface
func (m *MockConsumerRepository) Create(consumer *domain.Consumer) error {
	args := m.Called(consumer)
	return args.Error(0)
}

func (m *MockConsumerRepository) FindByID(id uint) (*domain.Consumer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Consumer), args.Error(1)
}

func (m *MockConsumerRepository) Update(consumer *domain.Consumer) error {
	args := m.Called(consumer)
	return args.Error(0)
}

func (m *MockConsumerRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Tambahan test untuk memastikan implementasi
func TestConsumerRepositoryImplementation(t *testing.T) {
	var repo repository.ConsumerRepository = &MockConsumerRepository{}
	assert.NotNil(t, repo)
}

func TestCreateConsumer(t *testing.T) {
	// Buat mock repository
	mockRepo := &MockConsumerRepository{}

	// Siapkan data konsumen untuk test
	consumer := &domain.Consumer{
		NIK:       "1234567890",
		FullName:  "John Doe",
		LegalName: "John Doe",
	}

	// Set expectation
	mockRepo.On("Create", consumer).Return(nil)

	// Buat service dengan mock repository
	consumerService := service.NewConsumerService(mockRepo)

	// Jalankan test
	err := consumerService.CreateConsumer(consumer)

	// Assertion
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
