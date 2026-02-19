package domain_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	params := domain.CreateUserParams{
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
	}
	user := domain.CreateUser(params)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, domain.StatusActive, user.Status)
	assert.Equal(t, domain.RoleUser, user.Role)
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

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "user@example.com", false},
		{"valid email with subdomain", "user@mail.example.com", false},
		{"valid email with plus", "user+tag@example.com", false},
		{"empty email", "", true},
		{"missing @", "userexample.com", true},
		{"missing domain", "user@", true},
		{"missing local part", "@example.com", true},
		{"spaces in email", "user @example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateEmail(tt.email)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, &domain.InvalidEmail{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"valid username", "johndoe", false},
		{"valid with numbers", "john123", false},
		{"valid with underscore", "john_doe", false},
		{"valid with hyphen", "john-doe", false},
		{"minimum length", "abc", false},
		{"too short", "ab", true},
		{"too long", "abcdefghijklmnopqrstuvwxyz1234567", true},
		{"starts with number", "1johndoe", true},
		{"starts with underscore", "_johndoe", true},
		{"contains space", "john doe", true},
		{"contains special char", "john@doe", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateUsername(tt.username)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, &domain.InvalidUsername{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{"empty phone (optional)", "", false},
		{"valid US format", "1234567890", false},
		{"valid with country code", "+1234567890", false},
		{"valid with dashes", "123-456-7890", false},
		{"valid with spaces", "123 456 7890", false},
		{"valid with parentheses", "(123) 456-7890", false},
		{"too short", "123456", true},
		{"contains letters", "123-ABC-7890", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidatePhone(tt.phone)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, &domain.InvalidPhone{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
