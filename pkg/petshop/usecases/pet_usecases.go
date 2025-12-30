// Package usecases provides use cases for petshop module.
package usecases

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
)

type PetUseCases struct {
	pr repositories.PetRepository
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
	Showher(string) (domain.Pet, error)
	Updater(string, PetUpdatersParams) (domain.Pet, error)
	Deleter(string) error
}

func NewPetUseCases(pr repositories.PetRepository) PetUseCases {
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

func (p *PetUseCases) Showher(petID string) (domain.Pet, error) {
	pet, err := p.pr.FindOne(petID)
	if err != nil {
		return domain.Pet{}, &domain.PetNotFound{ID: petID}
	}
	return *pet, nil
}

func (p *PetUseCases) Updater(petID string, payload PetUpdatersParams) (domain.Pet, error) {
	pet, err := p.pr.FindOne(petID)
	if err != nil {
		return domain.Pet{}, &domain.PetNotFound{ID: petID}
	}
	pet.Update(domain.UpdatePetParams(payload))

	err = p.pr.Save(*pet)
	if err != nil {
		return domain.Pet{}, err
	}

	return *pet, nil
}

func (p *PetUseCases) Deleter(petID string) error {
	_, err := p.pr.FindOne(petID)
	if err != nil {
		return &domain.PetNotFound{ID: petID}
	}
	return p.pr.Delete(petID)
}
