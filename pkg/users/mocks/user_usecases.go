package mocks

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/usecases"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCases struct {
	mock.Mock
}

var _ usecases.UserUseCasesInterface = &MockUserUseCases{}

func (m *MockUserUseCases) Creator(params usecases.UserCreatorParams) (domain.User, error) {
	args := m.Called(params)
	if args.Get(0) != nil {
		return args.Get(0).(domain.User), args.Error(1)
	}
	return domain.User{}, args.Error(1)
}

func (m *MockUserUseCases) Shower(username string) (domain.User, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(domain.User), args.Error(1)
	}
	return domain.User{}, args.Error(1)
}

func (m *MockUserUseCases) Updater(username string, params usecases.UserUpdatersParams) (domain.User, error) {
	args := m.Called(username, params)
	if args.Get(0) != nil {
		return args.Get(0).(domain.User), args.Error(1)
	}
	return domain.User{}, args.Error(1)
}

func (m *MockUserUseCases) Deleter(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserUseCases) LoggerIn(username string, password string) (string, string, error) {
	args := m.Called(username, password)
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.String(0), args.String(1), args.Error(2)
	}
	return "", "", args.Error(2)
}

func (m *MockUserUseCases) RefreshToken(refreshToken string) (string, string, error) {
	args := m.Called(refreshToken)
	if args.Get(0) != nil && args.Get(1) != nil {
		return args.String(0), args.String(1), args.Error(2)
	}
	return "", "", args.Error(2)
}

func (m *MockUserUseCases) LoggerOut(username string) error {
	args := m.Called(username)
	return args.Error(0)
}
