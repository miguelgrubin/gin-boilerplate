package application_test

import (
	"errors"
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"github.com/stretchr/testify/assert"
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
