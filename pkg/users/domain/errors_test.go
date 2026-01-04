package domain_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserNotFound(t *testing.T) {
	err := &domain.UserNotFound{ID: "1234"}
	expectedMessage := "User not found with ID: 1234"
	assert.Equal(t, expectedMessage, err.Error())
}

func TestUsernameNotFound(t *testing.T) {
	err := &domain.UsernameNotFound{Username: "johndoe"}
	expectedMessage := "User not found with Username: johndoe"
	assert.Equal(t, expectedMessage, err.Error())
}

func TestInvalidLogin(t *testing.T) {
	err := &domain.InvalidLogin{}
	expectedMessage := "Invalid login credentials"
	assert.Equal(t, expectedMessage, err.Error())
}
