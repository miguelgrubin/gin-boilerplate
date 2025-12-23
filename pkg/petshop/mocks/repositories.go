// Package mocks for pet repositories.
package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/stretchr/testify/mock"
)

type MockPetRepository struct {
	mock.Mock
}

func (m MockPetRepository) Save(pet domain.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func (m MockPetRepository) FindOne(id string) (*domain.Pet, error) {
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

func (m MockPetRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ repositories.PetRepository = &MockPetRepository{}
