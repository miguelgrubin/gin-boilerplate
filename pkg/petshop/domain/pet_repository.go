package domain

import "github.com/miguelgrubin/gin-boilerplate/pkg/shared"

type PetRepository interface {
	FindOne(shared.EntityID) (*Pet, error)
	FindAll() ([]Pet, error)
	Save(Pet) error
	Delete(shared.EntityID) error
}
