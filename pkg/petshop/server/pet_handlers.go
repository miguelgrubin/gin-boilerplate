package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
)

type PetHandlers struct {
	usecase usecases.PetUseCasesInterface
}

func NewPetHandlers(u usecases.PetUseCasesInterface) PetHandlers {
	return PetHandlers{
		usecase: u,
	}
}

// PetCreateRequest is the request payload for creating a new pet
func (ph *PetHandlers) PetCreateHandler(c *gin.Context) {
	var petParams PetCreateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pet, err := ph.usecase.Creator(usecases.PetCreatorParams{
		Name:   petParams.Name,
		Status: petParams.Status,
	})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
}

func (ph *PetHandlers) PetListHandler(c *gin.Context) {
	pets, err := ph.usecase.Finder(usecases.PetFinderParams{})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
}

func (ph *PetHandlers) PetShowHandler(c *gin.Context) {
	petID := c.Param("id")
	pet, err := ph.usecase.Showher(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (ph *PetHandlers) PetUpdateHandler(c *gin.Context) {
	petID := c.Param("id")
	var petParams PetUpdateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload := (usecases.PetUpdatersParams)(petParams)
	pet, err := ph.usecase.Updater(petID, payload)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (ph *PetHandlers) PetDeleteHandler(c *gin.Context) {
	petID := c.Param("id")
	err := ph.usecase.Deleter(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ph *PetHandlers) SetupRoutes(r *gin.RouterGroup) {
	r.POST("/pets", ph.PetCreateHandler)
	r.GET("/pets", ph.PetListHandler)
	r.GET("/pet/:id", ph.PetShowHandler)
	r.PATCH("/pet/:id", ph.PetUpdateHandler)
	r.DELETE("/pet/:id", ph.PetDeleteHandler)
}
