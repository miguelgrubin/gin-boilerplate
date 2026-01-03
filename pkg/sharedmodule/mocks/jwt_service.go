package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/mock"
)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateTokens(userID string, role string) (string, string, error) {
	args := m.Called(userID, role)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockJWTService) ValidateToken(token string) bool {
	args := m.Called(token)
	return args.Bool(0)
}

func (m *MockJWTService) RefreshToken(token string, userID string, role string) (string, string, error) {
	args := m.Called(token, userID, role)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockJWTService) DecodeToken(token string) (services.JWTData, error) {
	args := m.Called(token)
	return args.Get(0).(services.JWTData), args.Error(1)
}

func (m *MockJWTService) InvalidateToken(token string) error {
	args := m.Called(token)
	return args.Error(0)
}

var _ services.JWTService = &MockJWTService{}
