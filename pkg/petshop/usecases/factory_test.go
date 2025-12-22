package usecases_test

import (
	"testing"

	psMocks "github.com/miguelgrubin/gin-boilerplate/pkg/petshop/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
)

func TestNewPetShopUseCases(t *testing.T) {
	// Just a simple test to cover the factory function
	pr := new(psMocks.MockPetRepository)
	repos := repositories.PetShopRepositories{
		Pet: pr,
	}
	useCases := usecases.NewPetShopUseCases(repos)

	if useCases.Pet == nil {
		t.Error("Expected Pet use cases to be initialized")
	}
}
