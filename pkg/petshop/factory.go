// Package petshop provides a module of petshop.
package petshop

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/use_cases"
	"gorm.io/gorm"
)

type PetShopRepositories struct {
	pet repositories.PetRepository
}

func NewPetShopRepositories(db *gorm.DB) PetShopRepositories {
	petRepository := repositories.NewPetRepository(db)
	return PetShopRepositories{
		pet: petRepository,
	}
}

type PetShopUseCases struct {
	pet use_cases.PetUseCasesInterface
}

func NewPetShopUseCases(r PetShopRepositories) PetShopUseCases {
	petUseCases := use_cases.NewPetUseCases(r.pet)
	return PetShopUseCases{
		pet: &petUseCases,
	}
}

type PetShop struct {
	useCases     PetShopUseCases
	repositories PetShopRepositories
}

func NewPetShop(db *gorm.DB) PetShop {
	repositories := NewPetShopRepositories(db)
	useCases := NewPetShopUseCases(repositories)
	return PetShop{
		useCases:     useCases,
		repositories: repositories,
	}
}

func NewPetShopServer(db *gorm.DB, r *gin.RouterGroup) {
	petRepository := repositories.NewPetRepository(db)
	petUseCases := use_cases.NewPetUseCases(petRepository)
	server.NewPetRouterGroup(r, &petUseCases)
}

func NewPetShopMigrator(db *gorm.DB) {
	err := repositories.Automigrate(db)

	if err != nil {
		log.Print(err)
	}
}

func NewPetShopSeeder(db *gorm.DB) {
	_, err := repositories.SeedPets(db)

	if err != nil {
		log.Print(err)
	}
}
