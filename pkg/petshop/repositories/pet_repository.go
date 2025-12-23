package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

type PetRepository interface {
	FindOne(sharedmodule.EntityID) (*domain.Pet, error)
	FindAll() ([]domain.Pet, error)
	Save(domain.Pet) error
	Delete(sharedmodule.EntityID) error
}
