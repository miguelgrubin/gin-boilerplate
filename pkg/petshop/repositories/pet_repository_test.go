package repositories_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"
)

type PetRepositoryTestSuite struct {
	suite.Suite
	seededPets      []domain.Pet
	databaseService services.DBService
	petRepository   repositories.PetRepository
}

func (suite *PetRepositoryTestSuite) SetupTestSuite() {
	ds := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "sqlite3",
		Address: ":memory:",
	})
	suite.databaseService = ds
	suite.databaseService.Connect()
	suite.petRepository = repositories.NewPetRepository(ds.DB)
	suite.petRepository.Automigrate()
	suite.seededPets, _ = suite.petRepository.Seed()
}

func (suite *PetRepositoryTestSuite) TearDownTestSuite() {
	suite.databaseService.Close()
}

func (suite *PetRepositoryTestSuite) CleanDatabase() {
	db := suite.databaseService.GetDB()
	db.Exec("DELETE FROM pets")
	suite.seededPets, _ = suite.petRepository.Seed()
}

func (suite *PetRepositoryTestSuite) TestSaveWithNewPet() {
	pet := domain.NewPet(domain.CreatePetParams{Name: "testy", Status: "sleeping"})
	err := suite.petRepository.Save(pet)

	storedPet, _ := suite.petRepository.FindOne(pet.ID)
	suite.Equal(pet.ID, storedPet.ID)
	suite.NoError(err)
}

func (suite *PetRepositoryTestSuite) TestSaveWithStoredPet() {
	suite.CleanDatabase()

	pet := suite.seededPets[0]
	pet.Name = "New Name"
	err := suite.petRepository.Save(pet)

	storedPet, _ := suite.petRepository.FindOne(pet.ID)
	suite.Equal(storedPet.Name, "New Name")
	suite.NoError(err)
}

func (suite *PetRepositoryTestSuite) TestFindOneWithResult() {
	pet := suite.seededPets[0]

	storedPet, prErr := suite.petRepository.FindOne(pet.ID)

	suite.Equal(pet.ID, storedPet.ID)
	suite.NoError(prErr)
}

func (suite *PetRepositoryTestSuite) TestFindOneWithoutResult() {
	_, err := suite.petRepository.FindOne("random-id")

	suite.ErrorContains(err, "Pet not found")
	suite.Contains(err.Error(), "random-id")
}

func (suite *PetRepositoryTestSuite) TestFindAll() {
	suite.CleanDatabase()

	storedPets, err := suite.petRepository.FindAll()

	suite.Len(storedPets, len(suite.seededPets))
	suite.NoError(err)
}

func (suite *PetRepositoryTestSuite) TestDelete() {
	suite.CleanDatabase()
	pet := suite.seededPets[0]

	err := suite.petRepository.Delete(pet.ID)

	suite.NoError(err)
}

func TestPetRepository(t *testing.T) {
	prts := new(PetRepositoryTestSuite)
	prts.SetupTestSuite()
	suite.Run(t, prts)
	prts.TearDownTestSuite()
}
