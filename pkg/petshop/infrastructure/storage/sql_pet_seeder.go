package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

func SeedPet(db *gorm.DB) (*domain.Pet, error) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Fluffy",
		Status: "happy",
	})

	err := db.Create(&pet).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func SeedPets(db *gorm.DB) ([]domain.Pet, error) {
	pets := []domain.Pet{
		domain.NewPet(domain.CreatePetParams{
			Name:   "Tommy",
			Status: "bored",
		}),
		domain.NewPet(domain.CreatePetParams{
			Name:   "Katty",
			Status: "sleeping",
		}),
	}
	for _, v := range pets {
		err := db.Create(v).Error
		if err != nil {
			return nil, err
		}
	}
	return pets, nil
}
