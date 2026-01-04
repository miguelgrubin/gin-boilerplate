package domain_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := domain.NewUser()
	assert.NotEmpty(t, user)
}

func TestCreateUser(t *testing.T) {
	params := domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	}
	user := domain.CreateUser(params)
	assert.Equal(t, params.Username, user.Username)
	assert.Equal(t, params.FirstName, user.FirstName)
	assert.Equal(t, params.LastName, user.LastName)
	assert.Equal(t, params.Email, user.Email)
}

func TestUpdateUser(t *testing.T) {
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})

	newFirstName := "Jane"
	newLastName := "Smith"
	newEmail := "janesmith@example.com"
	newPhone := "1234567890"
	newStatus := "active"
	newPassword := "newpasswordhash"
	user.Update(domain.UpdateUserParams{
		FirstName: &newFirstName,
		LastName:  &newLastName,
		Email:     &newEmail,
		Phone:     &newPhone,
		Status:    &newStatus,
		Password:  &newPassword,
	})

	assert.Equal(t, "Jane", user.FirstName)
	assert.Equal(t, "Smith", user.LastName)
	assert.Equal(t, "janesmith@example.com", user.Email)
	assert.Equal(t, "1234567890", user.Phone)
	assert.Equal(t, "active", user.Status)
	assert.Equal(t, "newpasswordhash", user.PasswordHash)
}

func TestUpdateUserWithEmptyValues(t *testing.T) {
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})

	user.Update(domain.UpdateUserParams{})

	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "johndoe@example.com", user.Email)
}
