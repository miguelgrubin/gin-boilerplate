package pkg

import (
	"log"

	"github.com/gin-gonic/gin"
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

func (a *App) Migrate() error {
	a.SharedServices.DBService.Connect()
	defer a.SharedServices.DBService.Close()

	a.PetShopModule.Automigrate()
	a.UsersModule.Automigrate()

	return nil
}

func (a *App) Seed() error {
	a.SharedServices.DBService.Connect()
	defer a.SharedServices.DBService.Close()

	a.PetShopModule.Seed()
	return nil
}

func (a *App) WriteConfig() error {
	return a.SharedServices.ConfigService.WriteConfig()
}

func (a *App) GenerateKeys() error {
	a.SharedServices.RSAService.GenerateKeyPair()
	return a.SharedServices.RSAService.Write()
}

func (a *App) RunServer() {
	address := a.SharedServices.ConfigService.GetConfig().Server.Address

	a.SharedServices.DBService.Connect()
	defer a.SharedServices.DBService.Close()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check!")
	})
	v1 := r.Group("/v1")

	a.PetShopModule.SetupRoutes(v1)
	a.UsersModule.SetupRoutes(v1)

	err := r.Run(address)
	if err != nil {
		log.Print(err)
	}
}
