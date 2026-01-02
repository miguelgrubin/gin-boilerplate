package mocks

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/mock"
)

type MockRedisService struct {
	mock.Mock
}

func (m *MockRedisService) Set(key string, value any, expiration time.Duration) error {
	args := m.Called(key, value, expiration)
	return args.Error(0)
}

func (m *MockRedisService) Get(key string) (string, error) {
	args := m.Called(key)
	if args.Get(0) != nil {
		return args.String(0), args.Error(1)
	}
	return "", args.Error(1)
}

func (m *MockRedisService) Has(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockRedisService) Del(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

var _ services.RedisService = &MockRedisService{}
