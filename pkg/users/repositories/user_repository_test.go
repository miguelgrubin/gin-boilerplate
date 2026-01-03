package repositories_test

import (
	"testing"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/repositories"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	seededUsers     []domain.User
	databaseService services.DBService
	userRepository  repositories.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTestSuite() {
	ds := services.NewDBServiceGorm(services.DatabaseConfig{
		Driver:  "sqlite3",
		Address: ":memory:",
	})
	suite.databaseService = ds
	suite.databaseService.Connect()
	suite.userRepository = repositories.NewUserRepository(ds.DB)
	suite.userRepository.Automigrate()
}

func (suite *UserRepositoryTestSuite) TearDownTestSuite() {
	suite.databaseService.Close()
}

func (suite *UserRepositoryTestSuite) CleanDatabase() {
	db := suite.databaseService.GetDB()
	db.Exec("DELETE FROM users")
}

func TestUserRepository(t *testing.T) {
	ts := new(UserRepositoryTestSuite)
	ts.SetupTestSuite()
	suite.Run(t, ts)
	ts.TearDownTestSuite()
}

func (suite *UserRepositoryTestSuite) TestSaveWithNewUser() {
	suite.CleanDatabase()
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})
	err := suite.userRepository.Save(user)

	storedUser, _ := suite.userRepository.FindOne(user.ID)
	suite.Equal(user.ID, storedUser.ID)
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestSaveWithStoredUser() {
	suite.CleanDatabase()
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})
	err := suite.userRepository.Save(user)

	newName := "New Name"
	user.Update(domain.UpdateUserParams{
		FirstName: &newName,
	})
	err = suite.userRepository.Save(user)

	storedUser, _ := suite.userRepository.FindOne(user.ID)
	suite.Equal(storedUser.FirstName, newName)
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestFindOneWithResult() {
	suite.CleanDatabase()
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})
	err := suite.userRepository.Save(user)

	storedUser, urErr := suite.userRepository.FindOne(user.ID)
	suite.Equal(user.ID, storedUser.ID)
	suite.NoError(urErr)
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestFindOneWithoutResult() {
	suite.CleanDatabase()
	_, err := suite.userRepository.FindOne("random-id")

	suite.ErrorContains(err, "User not found")
	suite.Contains(err.Error(), "random-id")
}

func (suite *UserRepositoryTestSuite) TestFindOneByUsernameWithResult() {
	suite.CleanDatabase()
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})
	err := suite.userRepository.Save(user)

	storedUser, urErr := suite.userRepository.FindOneByUsername(user.Username)
	suite.Equal(user.Username, storedUser.Username)
	suite.NoError(urErr)
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestFindOneByUsernameWithoutResult() {
	suite.CleanDatabase()
	_, err := suite.userRepository.FindOneByUsername("random-username")

	suite.ErrorContains(err, "User not found with Username")
	suite.Contains(err.Error(), "random-username")
}

func (suite *UserRepositoryTestSuite) TestDelete() {
	suite.CleanDatabase()
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	})
	err := suite.userRepository.Save(user)

	err = suite.userRepository.Delete(user.ID)

	_, errNotFound := suite.userRepository.FindOne(user.ID)
	suite.ErrorContains(errNotFound, "User not found")
	suite.NoError(err)
}
