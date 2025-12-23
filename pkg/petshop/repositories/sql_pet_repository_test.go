package repositories_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/stretchr/testify/assert"
)

func TestSaveWithNewPet(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error("Connection error")
	}

	pet := domain.NewPet(domain.CreatePetParams{Name: "testy", Status: "sleeping"})
	pr := repositories.NewPetRepository(db)
	err = pr.Save(pet)

	storedPet, _ := pr.FindOne(pet.ID)
	assert.Equal(t, pet.ID.String(), storedPet.ID.String())
	assert.NoError(t, err)
}

func TestSaveWithStoredPet(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}
	pets, _ := repositories.SeedPets(db)

	pet := pets[0]
	pet.Name = "New Name"
	pr := repositories.NewPetRepository(db)
	err = pr.Save(pet)

	storedPet, _ := pr.FindOne(pet.ID)
	assert.Equal(t, storedPet.Name, "New Name")
	assert.NoError(t, err)
}

func TestFindOneWithResult(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}
	pets, _ := repositories.SeedPets(db)
	pet := pets[0]

	pr := repositories.NewPetRepository(db)
	storedPet, prErr := pr.FindOne(pet.ID)

	assert.Equal(t, pet.ID, storedPet.ID)
	assert.NoError(t, prErr)
}

func TestFindOneWithoutResult(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}

	pr := repositories.NewPetRepository(db)
	_, err = pr.FindOne(sharedmodule.EntityID("random-id"))

	assert.ErrorContains(t, err, "Pet not found")
	assert.Contains(t, err.Error(), "random-id")
}

func TestFindAllWithEmptyResult(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}

	pr := repositories.NewPetRepository(db)
	storedPets, err := pr.FindAll()

	assert.Empty(t, storedPets)
	assert.NoError(t, err)
}

func TestFindOneWithResults(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}
	pets, _ := repositories.SeedPets(db)

	pr := repositories.NewPetRepository(db)
	storedPets, err := pr.FindAll()

	assert.Len(t, storedPets, len(pets))
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error(err.Error())
	}
	pets, _ := repositories.SeedPets(db)
	pet := pets[0]

	pr := repositories.NewPetRepository(db)
	err = pr.Delete(pet.ID)

	assert.NoError(t, err)
}
