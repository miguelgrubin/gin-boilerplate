package petshop_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
	"github.com/stretchr/testify/assert"
)

func TestNewPetShopModule(t *testing.T) {
	db := sharedmodule.NewDbConnection("sqlite3", ":memory:")
	psModule := petshop.NewPetShopModule(db)

	assert.NotNil(t, psModule.Repositories, "Expected Repositories to be initialized")
	assert.NotNil(t, psModule.UseCases, "Expected UseCases to be initialized")
	assert.NotNil(t, psModule.Handlers, "Expected Handlers to be initialized")
}

func TestPetShopModuleAutomigrate(t *testing.T) {
	db := sharedmodule.NewDbConnection("sqlite3", ":memory:")
	psModule := petshop.NewPetShopModule(db)

	err := psModule.Automigrate(db)
	assert.Nil(t, err, "Expected Automigrate to complete without error")
}

func TestPetShopModuleSeed(t *testing.T) {
	db := sharedmodule.NewDbConnection("sqlite3", ":memory:")
	psModule := petshop.NewPetShopModule(db)

	err := psModule.Automigrate(db)
	if err != nil {
		t.Fatalf("Automigrate failed: %v", err)
	}

	err = psModule.Seed(db)
	assert.Nil(t, err, "Expected Seed to complete without error")
}
