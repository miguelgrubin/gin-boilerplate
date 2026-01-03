package domain

import (
	"time"

	"github.com/google/uuid"
	sd "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
)

const (
	UserUpdated = "user.updated"
	UserCreated = "user.created"
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

func NewUser() User {
	user := User{
		ID:            "",
		Username:      "",
		FirstName:     "",
		LastName:      "",
		Email:         "",
		Phone:         "",
		Status:        "",
		Role:          "",
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
		DeletedAt:     nil,
		eventRegistry: sd.NewEventRegistry(),
	}
	user.eventRegistry.AddEvent(UserCreated)
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
		Status:        "active",
		Role:          "user",
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
