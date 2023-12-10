package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

// PetCreateRequest is the request payload for creating a new pet
func (pc *PetShopController) PetCreateHandler(c *gin.Context) {
	var petParams PetCreateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	pet, err := pc.UseCases.Pet.Creator(usecases.PetCreatorParams{
		Name:   petParams.Name,
		Status: petParams.Status,
	})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, PetReponseFromDomain(pet))
}

// PetListHandler godoc
// @Summary      List all pets
// @Description  get string by ID
// @Tags         pets
// @Accept       json
// @Produce      json
// @Success      200  {array}  PetResponse
// @Router       /pets [get]
func (pc *PetShopController) PetListHandler(c *gin.Context) {
	pets, err := pc.UseCases.Pet.Finder(usecases.PetFinderParams{})

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetResponseListFromDomain(pets))
}

func (pc *PetShopController) PetShowHandler(c *gin.Context) {
	petID := shared.EntityID(c.Param("id"))
	pet, err := pc.UseCases.Pet.Showher(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (pc *PetShopController) PetUpdateHandler(c *gin.Context) {
	petID := shared.EntityID(c.Param("id"))
	var petParams PetUpdateRequest
	if err := c.ShouldBindJSON(&petParams); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload := (usecases.PetUpdatersParams)(petParams)
	pet, err := pc.UseCases.Pet.Updater(petID, payload)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, PetReponseFromDomain(pet))
}

func (pc *PetShopController) PetDeleteHandler(c *gin.Context) {
	petID := shared.EntityID(c.Param("id"))
	err := pc.UseCases.Pet.Deleter(petID)

	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
