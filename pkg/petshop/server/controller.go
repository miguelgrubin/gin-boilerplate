package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
)

type PetShopController struct {
	Repositories repositories.PetShopRepositories
	UseCases     usecases.PetShopUseCases
}

func NewPetShopController(r repositories.PetShopRepositories, u usecases.PetShopUseCases) PetShopController {
	return PetShopController{
		Repositories: r,
		UseCases:     u,
	}
}

func (p *PetShopController) SetupRoutes(r *gin.RouterGroup) {
	{
		r.POST("/pets", p.PetCreateHandler)
		r.GET("/pets", p.PetListHandler)
		r.GET("/pet/:id", p.PetShowHandler)
		r.PATCH("/pet/:id", p.PetUpdateHandler)
		r.DELETE("/pet/:id", p.PetDeleteHandler)
	}
}
