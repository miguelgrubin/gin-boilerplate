package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/stretchr/testify/mock"
)

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

func (m *MockPetUseCases) Showher(id sharedmodule.EntityID) (domain.Pet, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Updater(id sharedmodule.EntityID, params usecases.PetUpdatersParams) (domain.Pet, error) {
	args := m.Called(id, params)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetUseCases) Deleter(id sharedmodule.EntityID) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ usecases.PetUseCasesInterface = &MockPetUseCases{}
