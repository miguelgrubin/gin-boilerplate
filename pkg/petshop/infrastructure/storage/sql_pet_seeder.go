package storage

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"gorm.io/gorm"
)

func SeedPets(db *gorm.DB) ([]domain.Pet, error) {
	pr := NewPetRepository(db)
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
		err := pr.Save(v)
		if err != nil {
			return nil, err
		}
	}
	return pets, nil
}
