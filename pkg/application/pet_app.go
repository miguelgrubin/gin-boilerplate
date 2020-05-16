package application

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/entity"
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/repository"
)

type petApp struct {
	pr repository.PetRepository
}

var _ PetAppInterface = &petApp{}

type PetAppInterface interface {
	SavePet(*entity.Pet) (*entity.Pet, map[string]string)
	GetAllPets() ([]entity.Pet, error)
	GetPet(uint64) (*entity.Pet, error)
	UpdatePet(*entity.Pet) (*entity.Pet, map[string]string)
	DeletePet(uint64) error
}

func (p *petApp) SavePet(pet *entity.Pet) (*entity.Pet, map[string]string) {
	return p.pr.SavePet(pet)
}

func (p *petApp) GetAllPets() ([]entity.Pet, error) {
	return p.pr.GetAllPets()
}

func (p *petApp) GetPet(petId uint64) (*entity.Pet, error) {
	return p.pr.GetPet(petId)
}

func (p *petApp) UpdatePet(pet *entity.Pet) (*entity.Pet, map[string]string) {
	return p.pr.UpdatePet(pet)
}

func (p *petApp) DeletePet(petId uint64) error {
	return p.pr.DeletePet(petId)
}
