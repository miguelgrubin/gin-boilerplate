// Package users provides the users module.
package users

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/usecases"
)

type UsersModule struct {
	SharedServices sharedmodule.SharedModuleServices
	Repositories   UsersModuleRepositories
	UseCases       UsersModuleUseCases
	Handlers       UsersModuleHandlers
}

type UsersModuleUseCases struct {
	User usecases.UserUseCasesInterface
}

type UsersModuleRepositories struct {
	User repositories.UserRepository
}

type UsersModuleHandlers struct {
	User server.UserHandlers
}

func NewUsersModule(s sharedmodule.SharedModuleServices) UsersModule {
	db := s.DBService.GetDB()

	userRepository := repositories.NewUserRepository(db)
	userUseCases := usecases.NewUserUseCases(userRepository, s.JWTService)
	userHandlers := server.NewUserHandlers(&userUseCases)

	return UsersModule{
		SharedServices: s,
		Repositories: UsersModuleRepositories{
			User: userRepository,
		},
		UseCases: UsersModuleUseCases{
			User: &userUseCases,
		},
		Handlers: UsersModuleHandlers{
			User: userHandlers,
		},
	}
}

func (um UsersModule) Automigrate() error {
	err := um.Repositories.User.Automigrate()
	return err
}

func (um UsersModule) SetupRoutes(r *gin.RouterGroup) {
	um.Handlers.User.SetupRoutes(r)
}
