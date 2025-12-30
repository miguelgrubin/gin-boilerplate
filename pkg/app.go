package pkg

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users"
)

type App struct {
	SharedServices sharedmodule.SharedModuleServices
	PetShopModule  petshop.PetShopModule
	UsersModule    users.UsersModule
}

func NewApp() (*App, error) {
	sharedServices := sharedmodule.NewSharedModuleServices()
	petShopModule := petshop.NewPetShopModule(sharedServices)
	usersModule := users.NewUsersModule(sharedServices)

	app := &App{
		SharedServices: sharedServices,
		PetShopModule:  petShopModule,
		UsersModule:    usersModule,
	}

	return app, nil
}
