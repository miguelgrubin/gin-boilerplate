package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
)

type PetShopHandlers struct {
	usecases usecases.PetShopUseCases
}

func NewPetShopHandlers(u usecases.PetShopUseCases) PetShopHandlers {
	return PetShopHandlers{
		usecases: u,
	}
}

// PetCreateRequest is the request payload for creating a new pet
func (ph *PetShopHandlers) PetCreateHandler(c *gin.Context) {
	var petParams PetCreateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pet, err := ph.usecases.Pet.Creator(usecases.PetCreatorParams{
		Name:   petParams.Name,
		Status: petParams.Status,
	})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
}

func (ph *PetShopHandlers) PetListHandler(c *gin.Context) {
	pets, err := ph.usecases.Pet.Finder(usecases.PetFinderParams{})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
}

func (ph *PetShopHandlers) PetShowHandler(c *gin.Context) {
	petID := c.Param("id")
	pet, err := ph.usecases.Pet.Showher(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (ph *PetShopHandlers) PetUpdateHandler(c *gin.Context) {
	petID := c.Param("id")
	var petParams PetUpdateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload := (usecases.PetUpdatersParams)(petParams)
	pet, err := ph.usecases.Pet.Updater(petID, payload)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (ph *PetShopHandlers) PetDeleteHandler(c *gin.Context) {
	petID := c.Param("id")
	err := ph.usecases.Pet.Deleter(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ph *PetShopHandlers) SetupRoutes(r *gin.RouterGroup) {
	{
		r.POST("/pets", ph.PetCreateHandler)
		r.GET("/pets", ph.PetListHandler)
		r.GET("/pet/:id", ph.PetShowHandler)
		r.PATCH("/pet/:id", ph.PetUpdateHandler)
		r.DELETE("/pet/:id", ph.PetDeleteHandler)
	}
}
