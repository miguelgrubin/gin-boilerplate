package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

func NewPetRouterGroup(v1 *gin.RouterGroup, useCases application.PetUseCasesInterface) {
	v1.POST("/pets", func(c *gin.Context) {
		var petParams PetCreateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		pet, err := useCases.Creator(application.PetCreatorParams{
			Name:   petParams.Name,
			Status: petParams.Status,
		})

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
	})

	v1.GET("/pets", func(c *gin.Context) {
		pets, err := useCases.Finder(application.PetFinderParams{})

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
	})

	v1.GET("/pet/:id", func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		pet, err := useCases.Showher(petID)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetReponseFromDomain(pet))
	})

	v1.PATCH("/pet/:id", func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		var petParams PetUpdateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		payload := (application.PetUpdatersParams)(petParams)
		pet, err := useCases.Updater(petID, payload)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetReponseFromDomain(pet))
	})

	v1.DELETE("/pet/:id", func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		err := useCases.Deleter(petID)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})
}
