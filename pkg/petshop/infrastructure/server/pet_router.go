package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

func NewRouterGroup(v1 *gin.RouterGroup, useCases application.PetUseCasesInterface) {
	v1.POST("/pets", func(c *gin.Context) {
		pet, _ := useCases.Creator(application.PetCreatorParams{
			Name:   c.PostForm("name"),
			Status: c.PostForm("status"),
		})
		c.JSON(201, pet)
	})

	v1.GET("/pets", func(c *gin.Context) {
		pets, _ := useCases.Finder(application.PetFinderParams{})
		c.JSON(200, pets)
	})

	v1.GET("/pet/:id", func(c *gin.Context) {
		petId := shared.EntityId(c.Param("id"))
		pet, err := useCases.Showher(petId)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(200, PetReponseFromDomain(pet))
	})

	v1.PATCH("/pet/:id", func(c *gin.Context) {
		petId := shared.EntityId(c.Param("id"))
		c.Request.ParseForm()
		payload := application.PetUpdatersParams{
			Name:   c.PostForm("name"),
			Status: c.PostForm("status"),
		}
		pet, err := useCases.Updater(petId, payload)

		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(200, PetReponseFromDomain(pet))
	})

	v1.DELETE("/pet/:id", func(c *gin.Context) {
		petId := shared.EntityId(c.Param("id"))
		err := useCases.Deleter(petId)

		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(204, nil)
	})
}
