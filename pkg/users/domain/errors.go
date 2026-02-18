// Package domain provides domain definition of users entities and errors.
package domain

import "fmt"

type UserNotFound struct {
	ID string
}

func (p *UserNotFound) Error() string {
	return fmt.Sprintf("User not found with ID: %s", p.ID)
}

type UsernameNotFound struct {
	Username string
}

func (p *UsernameNotFound) Error() string {
	return fmt.Sprintf("User not found with Username: %s", p.Username)
}

type InvalidLogin struct{}

func (p *InvalidLogin) Error() string {
	return "Invalid login credentials"
}

// InvalidEmail represents an email validation error.
type InvalidEmail struct {
	Email string
}

func (p *InvalidEmail) Error() string {
	return fmt.Sprintf("Invalid email format: %s", p.Email)
}

// InvalidUsername represents a username validation error.
type InvalidUsername struct {
	Username string
	Reason   string
}

func (p *InvalidUsername) Error() string {
	return fmt.Sprintf("Invalid username '%s': %s", p.Username, p.Reason)
}

// InvalidPhone represents a phone validation error.
type InvalidPhone struct {
	Phone string
}

func (p *InvalidPhone) Error() string {
	return fmt.Sprintf("Invalid phone format: %s", p.Phone)
}
