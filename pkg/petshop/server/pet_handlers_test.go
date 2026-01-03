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
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	psMocks "github.com/miguelgrubin/gin-boilerplate/pkg/petshop/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createServerFixture(useCases usecases.PetUseCasesInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	os.Setenv("APP_ENV", "test")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	pc := server.NewPetHandlers(useCases)
	pc.SetupRoutes(v1)
	return router
}

func TestGetPets(t *testing.T) {
	pet := domain.NewPet()
	pet.ID = "pet-id"
	pet.Name = "Piggie"
	pet.Status = "Active"
	pets := []domain.Pet{pet}
	puc := new(psMocks.MockPetUseCases)
	puc.On("Finder", mock.AnythingOfType("usecases.PetFinderParams")).Return(pets, nil)
	router := createServerFixture(puc)
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
	puc := new(psMocks.MockPetUseCases)
	puc.On("Finder", mock.AnythingOfType("usecases.PetFinderParams")).Return(pets, errors.New("random error"))
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPet(t *testing.T) {
	pet := domain.NewPet()
	pet.ID = "pet-id"
	pet.Name = "Piggie"
	pet.Status = "Active"
	puc := new(psMocks.MockPetUseCases)
	puc.On("Showher", pet.ID).Return(pet, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", pet.ID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	var responseData server.PetResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseData)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, responseData.ID, pet.ID)
}

func TestGetPetWithNotFoundError(t *testing.T) {
	petID := "random-id"

	puc := new(psMocks.MockPetUseCases)
	puc.On("Showher", petID).Return(domain.Pet{}, &domain.PetNotFound{ID: petID})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePet(t *testing.T) {
	petID := "random-id"

	puc := new(psMocks.MockPetUseCases)
	puc.On("Deleter", petID).Return(nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petID)
	req, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeletePetWithError(t *testing.T) {
	petID := "random-id"

	puc := new(psMocks.MockPetUseCases)
	puc.On("Deleter", petID).Return(&domain.PetNotFound{ID: petID})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petID)
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
	pet := domain.CreatePet(domain.CreatePetParams{Name: name, Status: status})
	puc := new(psMocks.MockPetUseCases)
	puc.On("Creator", mock.AnythingOfType("usecases.PetCreatorParams")).Return(pet, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/pets", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostPetWithRequestError(t *testing.T) {
	invalidPayload := "{\"invalidKey\": false}"

	puc := new(psMocks.MockPetUseCases)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/pets", bytes.NewBufferString(invalidPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPatchPet(t *testing.T) {
	pet := domain.NewPet()
	pet.ID = "pet-id"
	pet.Name = "Pet Name"
	pet.Status = "Active"
	validPayload := "{\"status\": \"sleeping\"}"

	puc := new(psMocks.MockPetUseCases)
	puc.On("Updater", pet.ID, mock.AnythingOfType("usecases.PetUpdatersParams")).Return(domain.Pet{}, &domain.PetNotFound{ID: pet.ID})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", pet.ID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(validPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPatchPetWithNotFoundError(t *testing.T) {
	pet := domain.NewPet()
	pet.ID = "pet-id"
	pet.Name = "Pet Name"
	pet.Status = "Active"
	validPayload := "{\"status\": \"sleeping\"}"

	puc := new(psMocks.MockPetUseCases)
	puc.On("Updater", pet.ID, mock.AnythingOfType("usecases.PetUpdatersParams")).Return(pet, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", pet.ID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(validPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPatchPetWithRequestError(t *testing.T) {
	petID := "random-id"
	invalidPayload := "{\"status\": false}"
	puc := new(psMocks.MockPetUseCases)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/pet/%s", petID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBufferString(invalidPayload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
