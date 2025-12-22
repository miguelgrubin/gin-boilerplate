package repositories

import "gorm.io/gorm"

type PetShopRepositories struct {
	Pet PetRepository
}

func NewPetShopRepositories(db *gorm.DB) PetShopRepositories {
	petRepository := NewPetRepository(db)
	return PetShopRepositories{
		Pet: petRepository,
	}
}
