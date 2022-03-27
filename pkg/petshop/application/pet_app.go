package application

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

type PetUseCases struct {
	pr domain.PetRepository
}

var _ PetUseCasesInterface = &PetUseCases{}

type PetCreatorParams struct {
	Name   string
	Status string
}

type PetFinderParams PetCreatorParams

type PetUpdatersParams struct {
	Name   *string
	Status *string
}

type PetUseCasesInterface interface {
	Creator(PetCreatorParams) (domain.Pet, error)
	Finder(PetFinderParams) ([]domain.Pet, error)
	Showher(shared.EntityId) (domain.Pet, error)
	Updater(shared.EntityId, PetUpdatersParams) (domain.Pet, error)
	Deleter(shared.EntityId) error
}

func NewPetUseCases(pr domain.PetRepository) PetUseCases {
	return PetUseCases{pr}
}

func (p *PetUseCases) Creator(params PetCreatorParams) (domain.Pet, error) {
	pet := domain.NewPet(domain.CreatePetParams(params))
	err := p.pr.Save(pet)
	if err != nil {
		return domain.Pet{}, err
	}
	return pet, nil
}

func (p *PetUseCases) Finder(_ PetFinderParams) ([]domain.Pet, error) {
	return p.pr.FindAll()
}

func (p *PetUseCases) Showher(petId shared.EntityId) (domain.Pet, error) {
	pet, err := p.pr.FindOne(petId)
	if err != nil {
		return domain.Pet{}, err
	}
	return *pet, nil
}

func (p *PetUseCases) Updater(petId shared.EntityId, payload PetUpdatersParams) (domain.Pet, error) {
	pet, err := p.pr.FindOne(petId)
	if err != nil {
		return domain.Pet{}, err
	}
	pet.Update(domain.UpdatePetParams(payload))

	err = p.pr.Save(*pet)
	if err != nil {
		return domain.Pet{}, err
	}

	return *pet, nil
}

func (p *PetUseCases) Deleter(petId shared.EntityId) error {
	return p.pr.Delete(petId)
}
