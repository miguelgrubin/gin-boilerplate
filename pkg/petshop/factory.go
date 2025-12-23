package petshop

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"gorm.io/gorm"
)

type PetShopModule struct {
	Repositories repositories.PetShopRepositories
	UseCases     usecases.PetShopUseCases
	Handlers     server.PetShopHandlers
}

func NewPetShopModule(db *gorm.DB) PetShopModule {
	psRepositories := repositories.NewPetShopRepositories(db)
	psUseCases := usecases.NewPetShopUseCases(psRepositories)
	psHandlers := server.NewPetShopHandlers(psUseCases)
	return PetShopModule{
		Repositories: psRepositories,
		UseCases:     psUseCases,
		Handlers:     psHandlers,
	}
}

func (ps PetShopModule) Automigrate(db *gorm.DB) error {
	return repositories.Automigrate(db)
}

func (ps PetShopModule) Seed(db *gorm.DB) error {
	_, err := repositories.SeedPets(db)
	return err
}
