package domain

import (
	"regexp"
	"time"
	"unicode"

	"github.com/google/uuid"
	sd "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
)

const (
	UserUpdated = "user.updated"
	UserCreated = "user.created"
)

// User status constants
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

// User role constants
const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID            string
	Username      string
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	PasswordHash  string
	Status        string
	Role          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	eventRegistry *sd.EventRegistry
}

type CreateUserParams struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
}

type UpdateUserParams struct {
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	Phone     *string
	Status    *string
}

// newUser creates a User with an initialized event registry.
// This is intended for internal use when hydrating entities from persistence.
// For creating new users, use CreateUser instead.
func newUser() User {
	return User{
		eventRegistry: sd.NewEventRegistry(),
	}
}

// HydrateUser creates a User from persistence data.
// This should only be used by repository mappers to reconstruct entities from storage.
// For creating new users in the application, use CreateUser instead.
func HydrateUser(
	id, username, firstName, lastName, email, phone, passwordHash, status, role string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) User {
	user := newUser()
	user.ID = id
	user.Username = username
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Phone = phone
	user.PasswordHash = passwordHash
	user.Status = status
	user.Role = role
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	user.DeletedAt = deletedAt
	return user
}

func CreateUser(payload CreateUserParams) User {
	id := uuid.New().String()
	user := User{
		ID:            id,
		Username:      payload.Username,
		FirstName:     payload.FirstName,
		LastName:      payload.LastName,
		Email:         payload.Email,
		Phone:         payload.Phone,
		Status:        StatusActive,
		Role:          RoleUser,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
		DeletedAt:     nil,
		eventRegistry: sd.NewEventRegistry(),
	}
	user.eventRegistry.AddEvent(UserCreated)
	return user
}

func (p *User) Update(payload UpdateUserParams) {
	if payload.FirstName != nil {
		p.FirstName = *payload.FirstName
	}
	if payload.LastName != nil {
		p.LastName = *payload.LastName
	}
	if payload.Email != nil {
		p.Email = *payload.Email
	}
	if payload.Password != nil {
		p.PasswordHash = *payload.Password
	}
	if payload.Phone != nil {
		p.Phone = *payload.Phone
	}
	if payload.Status != nil {
		p.Status = *payload.Status
	}
	p.UpdatedAt = time.Now()
	p.eventRegistry.AddEvent(UserUpdated)
}

// Username validation constraints
const (
	UsernameMinLength = 3
	UsernameMaxLength = 32
)

// Email validation regex pattern (RFC 5322 simplified)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Phone validation regex pattern (allows digits, spaces, dashes, parentheses, and optional + prefix)
var phoneRegex = regexp.MustCompile(`^\+?[\d\s\-()]{7,20}$`)

// ValidateEmail checks if the provided email has a valid format.
func ValidateEmail(email string) error {
	if email == "" {
		return &InvalidEmail{Email: email}
	}
	if !emailRegex.MatchString(email) {
		return &InvalidEmail{Email: email}
	}
	return nil
}

// ValidateUsername checks if the provided username meets the requirements:
// - Length between 3 and 32 characters
// - Contains only alphanumeric characters, underscores, and hyphens
// - Starts with a letter
func ValidateUsername(username string) error {
	if len(username) < UsernameMinLength {
		return &InvalidUsername{
			Username: username,
			Reason:   "must be at least 3 characters long",
		}
	}
	if len(username) > UsernameMaxLength {
		return &InvalidUsername{
			Username: username,
			Reason:   "must be at most 32 characters long",
		}
	}

	// Check first character is a letter
	runes := []rune(username)
	if !unicode.IsLetter(runes[0]) {
		return &InvalidUsername{
			Username: username,
			Reason:   "must start with a letter",
		}
	}

	// Check all characters are valid (letters, digits, underscores, hyphens)
	for _, r := range runes {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' && r != '-' {
			return &InvalidUsername{
				Username: username,
				Reason:   "can only contain letters, digits, underscores, and hyphens",
			}
		}
	}

	return nil
}

// ValidatePhone checks if the provided phone number has a valid format.
// Accepts digits, spaces, dashes, parentheses, and optional + prefix.
// Empty phone is considered valid (optional field).
func ValidatePhone(phone string) error {
	if phone == "" {
		return nil // Phone is optional
	}
	if !phoneRegex.MatchString(phone) {
		return &InvalidPhone{Phone: phone}
	}
	return nil
}
