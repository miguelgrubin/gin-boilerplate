package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	repositoryMocks "github.com/miguelgrubin/gin-boilerplate/mocks/pkg/petshop/repositories"
	mocks "github.com/miguelgrubin/gin-boilerplate/mocks/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createServerFixture(t *testing.T, useCases usecases.PetUseCasesInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	os.Setenv("APP_ENV", "test")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	petShopRepositories := repositories.PetShopRepositories{
		Pet: new(repositoryMocks.PetRepository),
	}
	petShopUseCases := usecases.PetShopUseCases{
		Pet: useCases,
	}
	pc := server.NewPetShopController(petShopRepositories, petShopUseCases)
	pc.SetupRoutes(v1)
	return router
}

func TestGetPets(t *testing.T) {
	pets := []domain.Pet{domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})}
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Finder", mock.AnythingOfType("usecases.PetFinderParams")).Return(pets, nil)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pets", nil)
	router.ServeHTTP(w, req)

	var responseData []server.PetResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, responseData)
}

func TestGetPetsWithRandomError(t *testing.T) {
	pets := []domain.Pet{}
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Finder", mock.AnythingOfType("usecases.PetFinderParams")).Return(pets, errors.New("random error"))
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPet(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{
		Name:   "Piggie",
		Status: "Active",
	})
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Showher", pet.ID).Return(pet, nil)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", pet.ID.String())
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	var responseData server.PetResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, responseData.ID, pet.ID.String())
}

func TestGetPetWithNotFoundError(t *testing.T) {
	petId := shared.EntityID("random-id")
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Showher", petId).Return(domain.Pet{}, &domain.PetNotFound{ID: petId.String()})
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePet(t *testing.T) {
	petId := shared.EntityID("random-id")
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Deleter", petId).Return(nil)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeletePetWithError(t *testing.T) {
	petId := shared.EntityID("random-id")
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Deleter", petId).Return(&domain.PetNotFound{ID: petId.String()})
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostPet(t *testing.T) {
	name := "Pet Name"
	status := "Active"
	body, err := json.Marshal(server.PetCreateRequest{
		Name:   name,
		Status: status,
	})
	pet := domain.NewPet(domain.CreatePetParams{Name: name, Status: status})
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Creator", mock.AnythingOfType("usecases.PetCreatorParams")).Return(pet, nil)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/pets", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostPetWithRequestError(t *testing.T) {
	invalidPayload := "{\"invalidKey\": false}"
	puc := new(mocks.PetUseCasesInterface)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/pets", bytes.NewBufferString(invalidPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPatchPet(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{Name: "Pet Name", Status: "Active"})
	petId := pet.ID
	validPayload := "{\"status\": \"sleeping\"}"
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Updater", petId, mock.AnythingOfType("usecases.PetUpdatersParams")).Return(domain.Pet{}, &domain.PetNotFound{ID: petId.String()})
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(validPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPatchPetWithNotFoundError(t *testing.T) {
	pet := domain.NewPet(domain.CreatePetParams{Name: "Pet Name", Status: "Active"})
	petId := pet.ID
	validPayload := "{\"status\": \"sleeping\"}"
	puc := new(mocks.PetUseCasesInterface)
	puc.On("Updater", petId, mock.AnythingOfType("usecases.PetUpdatersParams")).Return(pet, nil)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(validPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPatchPetWithRequestError(t *testing.T) {
	petId := shared.EntityID("random-id")
	invalidPayload := "{\"status\": false}"
	puc := new(mocks.PetUseCasesInterface)
	router := createServerFixture(t, puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petId.String())
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(invalidPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
