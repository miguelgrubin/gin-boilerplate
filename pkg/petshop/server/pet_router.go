package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/use_cases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

func PetCreateHandler(uc use_cases.PetUseCasesInterface) gin.HandlerFunc {
	// PetCreateRequest is the request payload for creating a new pet
	return func(c *gin.Context) {
		var petParams PetCreateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		pet, err := uc.Creator(use_cases.PetCreatorParams{
			Name:   petParams.Name,
			Status: petParams.Status,
		})

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
	}
}

func PetListHandler(uc use_cases.PetUseCasesInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		pets, err := uc.Finder(use_cases.PetFinderParams{})

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
	}
}

func PetShowHandler(uc use_cases.PetUseCasesInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		pet, err := uc.Showher(petID)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetReponseFromDomain(pet))
	}
}

func PetUpdateHandler(uc use_cases.PetUseCasesInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		var petParams PetUpdateRequest
		if err := c.ShouldBindJSON(&petParams); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		payload := (use_cases.PetUpdatersParams)(petParams)
		pet, err := uc.Updater(petID, payload)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, PetReponseFromDomain(pet))
	}
}

func PetDeleteHandler(uc use_cases.PetUseCasesInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		petID := shared.EntityID(c.Param("id"))
		err := uc.Deleter(petID)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

func NewPetRouterGroup(v1 *gin.RouterGroup, useCases use_cases.PetUseCasesInterface) {
	v1.POST("/pets", PetCreateHandler(useCases))
	v1.GET("/pets", PetListHandler(useCases))
	v1.GET("/pet/:id", PetShowHandler(useCases))
	v1.PATCH("/pet/:id", PetUpdateHandler(useCases))
	v1.DELETE("/pet/:id", PetDeleteHandler(useCases))
}
