package domain_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewPet(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	assert.Equal(t, pet.Name, "Piggie")
	assert.Equal(t, pet.Status, "Active")
}

func TestPetUpdateName(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	newName := "Peggie"
	pet.Update(domain.UpdatePetParams{
		Name: &newName,
	})
	assert.Equal(t, pet.Name, newName)
	assert.Equal(t, pet.Status, "Active")
}

func TestPetUpdateStatus(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	newStatus := "Sleeping"
	pet.Update(domain.UpdatePetParams{
		Status: &newStatus,
	})
	assert.Equal(t, pet.Name, "Piggie")
	assert.Equal(t, pet.Status, newStatus)
}
