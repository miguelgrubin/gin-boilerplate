package application_test

import (
	"errors"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPetShowerWhenHasResult(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(&pet, nil)
	useCases := application.NewPetUseCases(pr)

	result, _ := useCases.Showher(shared.EntityID(pet.ID))

	pr.AssertExpectations(t)
	assert.Equal(t, result, pet)
}

func TestPetShowerWhenHasNotResult(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})

	prError := errors.New("Random error from db layer")
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(nil, prError)
	useCases := application.NewPetUseCases(pr)
	domainErr := &domain.PetNotFound{ID: pet.ID.AsString()}

	_, err := useCases.Showher(shared.EntityID(pet.ID))
	pr.AssertExpectations(t)
	assert.Equal(t, err, domainErr)
}

func TestPetCreatorWithNoErrorOnSave(t *testing.T) {
	const name = "Piggie"
	const status = "Active"

	pr := new(mocks.PetRepository)
	pr.On("Save", mock.AnythingOfType("domain.Pet")).Return(nil)
	useCases := application.NewPetUseCases(pr)
	pet, err := useCases.Creator(application.PetCreatorParams{
		Name:   name,
		Status: status,
	})

	pr.AssertExpectations(t)
	assert.Equal(t, pet.Name, name)
	assert.Equal(t, pet.Status, status)
	assert.NoError(t, err)
}

func TestPetCreatorWithErrorOnSave(t *testing.T) {
	const name = "Piggie"
	const status = "Active"

	pr := new(mocks.PetRepository)
	pr.On("Save", mock.AnythingOfType("domain.Pet")).Return(errors.New("generic error from repository"))
	useCases := application.NewPetUseCases(pr)
	_, err := useCases.Creator(application.PetCreatorParams{
		Name:   name,
		Status: status,
	})

	pr.AssertExpectations(t)
	assert.ErrorContains(t, err, "error from repository")
}

func TestPetFinder(t *testing.T) {
	var pets []domain.Pet = make([]domain.Pet, 0)
	pr := new(mocks.PetRepository)
	pr.On("FindAll").Return(pets, nil)

	useCases := application.NewPetUseCases(pr)
	result, err := useCases.Finder(application.PetFinderParams{})

	pr.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, result, pets)
}

func TestPetUpdaterWithExistantPet(t *testing.T) {
	newName := "New Name"
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(&pet, nil)
	pr.On("Save", mock.AnythingOfType("domain.Pet")).Return(nil)

	useCases := application.NewPetUseCases(pr)
	result, err := useCases.Updater(pet.ID, application.PetUpdatersParams{Name: &newName})

	pr.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, result.Name, newName)
}

func TestPetUpdaterWithUnexistantPet(t *testing.T) {
	newName := "New Name"
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(&domain.Pet{}, &domain.PetNotFound{ID: pet.ID.AsString()})

	useCases := application.NewPetUseCases(pr)
	_, err := useCases.Updater(pet.ID, application.PetUpdatersParams{Name: &newName})

	pr.AssertExpectations(t)
	assert.ErrorContains(t, err, pet.ID.AsString())
}

func TestPetDeleterWithExistantPet(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(&pet, nil)
	pr.On("Delete", pet.ID).Return(nil)

	useCases := application.NewPetUseCases(pr)
	err := useCases.Deleter(pet.ID)

	pr.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestPetDeleterWithUnexistantPet(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	pr := new(mocks.PetRepository)
	pr.On("FindOne", pet.ID).Return(&domain.Pet{}, &domain.PetNotFound{ID: pet.ID.AsString()})

	useCases := application.NewPetUseCases(pr)
	err := useCases.Deleter(pet.ID)

	pr.AssertExpectations(t)
	assert.ErrorContains(t, err, pet.ID.AsString())
}
