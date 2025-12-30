// Package domain provides domain definition of petshop entities and errors.
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
