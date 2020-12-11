package repository

import "github.com/miguelgrubin/gin-boilerplate/pkg/domain/entity"

type PetRepository interface {
	SavePet(*entity.Pet) (*entity.Pet, map[string]string)
	GetPet(uint64) (*entity.Pet, error)
	GetAllPets() ([]entity.Pet, error)
	UpdatePet(*entity.Pet) (*entity.Pet, map[string]string)
	DeletePet(uint64) error
}
