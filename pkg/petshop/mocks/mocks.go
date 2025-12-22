package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"github.com/stretchr/testify/mock"
)

type MockPetRepository struct {
	mock.Mock
}

func (m MockPetRepository) Save(pet domain.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func (m MockPetRepository) FindOne(id shared.EntityID) (*domain.Pet, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Pet), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m MockPetRepository) FindAll() ([]domain.Pet, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Pet), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m MockPetRepository) Delete(id shared.EntityID) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ repositories.PetRepository = &MockPetRepository{}

type MockPetUseCases struct {
	mock.Mock
}

func (m *MockPetUseCases) Creator(params usecases.PetCreatorParams) (domain.Pet, error) {
	args := m.Called(params)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Finder(params usecases.PetFinderParams) ([]domain.Pet, error) {
	args := m.Called(params)
	return args.Get(0).([]domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Showher(id shared.EntityID) (domain.Pet, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Updater(id shared.EntityID, params usecases.PetUpdatersParams) (domain.Pet, error) {
	args := m.Called(id, params)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Deleter(id shared.EntityID) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ usecases.PetUseCasesInterface = &MockPetUseCases{}
