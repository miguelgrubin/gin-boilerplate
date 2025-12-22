package repositories_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNewPetRepository(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Error("Connection error")
	}

	psr := repositories.NewPetShopRepositories(db)

	assert.NotNil(t, psr.Pet, "Expected Pet repository to be initialized")
	assert.IsType(t, repositories.SQLPetRepository{}, psr.Pet, "Expected Pet repository to be of type SQLPetRepository")
}
