package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/middlewares"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
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
	err := a.SharedServices.DBService.Connect()
	if err != nil {
		a.SharedServices.LoggerService.Error("error connecting to database", services.Err(err))
		return err
	}
	defer a.SharedServices.DBService.Close()

	err = a.PetShopModule.Automigrate()
	if err != nil {
		return err
	}
	err = a.UsersModule.Automigrate()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Seed() error {
	defer a.SharedServices.DBService.Close()

	err := a.PetShopModule.Seed()
	return err
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
	logger := a.SharedServices.LoggerService

	defer a.SharedServices.DBService.Close()
	defer func() { _ = logger.Sync() }()

	r := gin.New()
	r.Use(middlewares.RequestLogger(logger))
	r.Use(gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check!")
	})
	v1 := r.Group("/v1")

	a.PetShopModule.SetupRoutes(v1)
	a.UsersModule.SetupRoutes(v1)

	logger.Info("starting server", services.String("address", address))
	err := r.Run(address)
	if err != nil {
		logger.Error("server error", services.Err(err))
	}
}
