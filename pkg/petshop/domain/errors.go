// Package domain provides domain definition of petshop entities and errors.
package domain

import "fmt"

type PetNotFound struct {
	ID string
}

func (p *PetNotFound) Error() string {
	return fmt.Sprintf("Pet not found with ID: %s", p.ID)
}
