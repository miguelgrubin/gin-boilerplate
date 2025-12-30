// Package petshop provides the petshop module.
package petshop

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

type PetShopModule struct {
	SharedServices sharedmodule.SharedModuleServices
	Repositories   PetShopModuleRepositories
	UseCases       PetShopModuleUseCases
	Handlers       PetShopModuleHandlers
}

type PetShopModuleUseCases struct {
	Pet usecases.PetUseCasesInterface
}

type PetShopModuleRepositories struct {
	Pet repositories.PetRepository
}

type PetShopModuleHandlers struct {
	Pet server.PetHandlers
}

func NewPetShopModule(s sharedmodule.SharedModuleServices) PetShopModule {
	db := s.DBService.GetDB()

	pr := repositories.NewPetRepository(db)
	pu := usecases.NewPetUseCases(pr)
	ph := server.NewPetHandlers(&pu)
	return PetShopModule{
		SharedServices: s,
		Repositories: PetShopModuleRepositories{
			Pet: pr,
		},
		UseCases: PetShopModuleUseCases{
			Pet: &pu,
		},
		Handlers: PetShopModuleHandlers{
			Pet: ph,
		},
	}
}

func (ps PetShopModule) SetupRoutes(r *gin.RouterGroup) {
	ps.Handlers.Pet.SetupRoutes(r)
}

func (ps PetShopModule) Automigrate() error {
	err := ps.Repositories.Pet.Automigrate()
	return err
}

func (ps PetShopModule) Seed() error {
	_, err := ps.Repositories.Pet.Seed()
	return err
}
