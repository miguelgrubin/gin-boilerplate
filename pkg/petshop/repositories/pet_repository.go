package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

type PetRepository interface {
	FindOne(shared.EntityID) (*domain.Pet, error)
	FindAll() ([]domain.Pet, error)
	Save(domain.Pet) error
	Delete(shared.EntityID) error
}
