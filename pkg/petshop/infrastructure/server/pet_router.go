package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

func NewRouterGroup(v1 *gin.RouterGroup, useCases application.PetUseCasesInterface) {
	v1.POST("/pets", func(c *gin.Context) {
		var petParams PetCreateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		pet, _ := useCases.Creator(application.PetCreatorParams{
			Name:   petParams.Name,
			Status: petParams.Status,
		})
		c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
	})

	v1.GET("/pets", func(c *gin.Context) {
		pets, _ := useCases.Finder(application.PetFinderParams{})
		c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
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
		var petParams PetUpdateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		payload := (application.PetUpdatersParams)(petParams)
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
