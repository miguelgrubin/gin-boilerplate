package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/repositories"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindOne(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindOneByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) Automigrate() error {
	args := m.Called()
	return args.Error(0)
}

var _ repositories.UserRepository = &MockUserRepository{}
