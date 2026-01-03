package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/mock"
)

type MockHashService struct {
	mock.Mock
}

func (m *MockHashService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockHashService) Compare(hash string, password string) bool {
	args := m.Called(hash, password)
	return args.Bool(0)
}

var _ services.HashService = &MockHashService{}
