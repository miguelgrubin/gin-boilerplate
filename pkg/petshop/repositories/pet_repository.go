package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

type PetRepository interface {
	FindOne(string) (*domain.Pet, error)
	FindAll() ([]domain.Pet, error)
	Save(domain.Pet) error
	Delete(string) error
}
