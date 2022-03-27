package domain

import "github.com/miguelgrubin/gin-boilerplate/pkg/shared"

type PetRepository interface {
	FindOne(shared.EntityId) (*Pet, error)
	FindAll() ([]Pet, error)
	Save(Pet) error
	Delete(shared.EntityId) error
}
