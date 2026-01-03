package usecases_test

import (
	"errors"
	"testing"

	sd "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
	mockServices "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecasesTestSuite struct {
	suite.Suite
	userRepositoryMock *mocks.MockUserRepository
	hashServiceMock    *mockServices.MockHashService
	jwtServiceMock     *mockServices.MockJWTService
}

func (ts *UserUsecasesTestSuite) TestCreator() {
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("hashedpassword", nil)
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	username := "mm"
	firstName := "Martín"
	lastName := "Martínez"
	email := "mm@example.com"
	password := "securepassword"
	phone := "1234567890"
	user, err := uc.Creator(usecases.UserCreatorParams{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Phone:     phone,
	})

	ts.NoError(err)
	ts.Equal(username, user.Username)
	ts.Equal(firstName, user.FirstName)
	ts.Equal(lastName, user.LastName)
	ts.Equal(email, user.Email)
	ts.Equal(phone, user.Phone)
	ts.NotEmpty(user.ID)
}

func (ts *UserUsecasesTestSuite) TestCreatorHashError() {
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("", errors.New("hash error"))
	ts.userRepositoryMock.On("Save", mock.Anything).Return(nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, err := uc.Creator(usecases.UserCreatorParams{
		Username:  "mm",
		FirstName: "Martín",
		LastName:  "Martínez",
		Email:     "mm@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})

	ts.Error(err)
}

func (ts *UserUsecasesTestSuite) TestCreatorSaveError() {
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("hashedpassword", nil)
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(errors.New("save error"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, err := uc.Creator(usecases.UserCreatorParams{
		Username:  "mm",
		FirstName: "Martín",
		LastName:  "Martínez",
		Email:     "mm@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})

	ts.Error(err)
}

func (ts *UserUsecasesTestSuite) TestShower() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", "user-id").Return(&user, nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	returnedUser, err := uc.Shower("user-id")

	ts.NoError(err)
	ts.Equal(user.ID, returnedUser.ID)
	ts.Equal(user.Username, returnedUser.Username)
}

func (ts *UserUsecasesTestSuite) TestShowerNotFound() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", "user-id").Return(nil, errors.New("user not found"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, err := uc.Shower("user-id")

	ts.Error(err)
	ts.IsType(&domain.UsernameNotFound{}, err)
}

func (ts *UserUsecasesTestSuite) TestUpdater() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(nil)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("newhashedpassword", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	fname := "Martín Martín"
	userResponse, err := uc.Updater("mm", usecases.UserUpdatersParams{
		FirstName: &fname,
	})

	ts.NoError(err)
	ts.Equal(fname, userResponse.FirstName)
}

func (ts *UserUsecasesTestSuite) TestUpdaterNotFound() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(nil, errors.New("user not found"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	pwd := "newsecurepassword"
	_, err := uc.Updater("mm", usecases.UserUpdatersParams{
		Password: &pwd,
	})

	ts.Error(err)
}

func (ts *UserUsecasesTestSuite) TestUpdaterWithPassword() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(nil)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("newhashedpassword", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	pwd := "newsecurepassword"
	userResponse, err := uc.Updater("mm", usecases.UserUpdatersParams{
		Password: &pwd,
	})

	ts.NoError(err)
	ts.Equal("newhashedpassword", userResponse.PasswordHash)
}

func (ts *UserUsecasesTestSuite) TestUpdaterWithErrorHashing() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(nil)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("", errors.New("hashing error"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	pwd := "newsecurepassword"
	_, err := uc.Updater("mm", usecases.UserUpdatersParams{
		Password: &pwd,
	})

	ts.Error(err)
}

func (ts *UserUsecasesTestSuite) TestUpdaterWithErrorSaving() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.userRepositoryMock.On("Save", mock.Anything).Return(errors.New("saving error"))
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Hash", mock.Anything).Return("newhashedpassword", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	pwd := "newsecurepassword"
	_, err := uc.Updater("mm", usecases.UserUpdatersParams{
		Password: &pwd,
	})

	ts.Error(err)
}

func (ts *UserUsecasesTestSuite) TestDeleter() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	user := mockedUser()
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.userRepositoryMock.On("Delete", mock.Anything).Return(nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	err := uc.Deleter("mm")

	ts.NoError(err)
}

func (ts *UserUsecasesTestSuite) TestDeleterNotFound() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(nil, errors.New("user not found"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	err := uc.Deleter("mm")

	ts.Error(err)
	ts.IsType(&domain.UsernameNotFound{}, err)
}

func (ts *UserUsecasesTestSuite) TestLoggerIn() {
	user := mockedUser()
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Compare", mock.Anything, mock.Anything).Return(true)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("GenerateTokens", mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	token, refreshToken, err := uc.LoggerIn("mm", "securepassword")

	ts.NoError(err)
	ts.NotEmpty(token)
	ts.NotEmpty(refreshToken)
}

func (ts *UserUsecasesTestSuite) TestLoggerInNotFound() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(nil, errors.New("user not found"))
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Compare", mock.Anything, mock.Anything).Return(true)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("GenerateTokens", mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, _, err := uc.LoggerIn("mm", "securepassword")

	ts.Error(err)
	ts.IsType(&domain.InvalidLogin{}, err)
}

func (ts *UserUsecasesTestSuite) TestLoggerInInvalidPassword() {
	user := mockedUser()
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOneByUsername", mock.Anything).Return(&user, nil)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.hashServiceMock.On("Compare", mock.Anything, mock.Anything).Return(false)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("GenerateTokens", mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, _, err := uc.LoggerIn("mm", "securepassword")

	ts.Error(err)
	ts.IsType(&domain.InvalidLogin{}, err)
}

func (ts *UserUsecasesTestSuite) TestRefreshToken() {
	user := mockedUser()
	tokenData := services.JWTData{
		UserID: user.ID,
		Role:   user.Role,
	}
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOne", mock.Anything).Return(&user, nil)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("DecodeToken", mock.Anything).Return(tokenData, nil)
	ts.jwtServiceMock.On("RefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	newToken, newRefreshToken, err := uc.RefreshToken("sometoken")

	ts.NoError(err)
	ts.NotEmpty(newToken)
	ts.NotEmpty(newRefreshToken)
}

func (ts *UserUsecasesTestSuite) TestRefreshTokenInvalidToken() {
	user := mockedUser()
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOne", mock.Anything).Return(&user, nil)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("DecodeToken", mock.Anything).Return(services.JWTData{}, errors.New("invalid token"))
	ts.jwtServiceMock.On("RefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, _, err := uc.RefreshToken("sometoken")

	ts.Error(err)
	ts.IsType(&sd.InvalidRefreshToken{}, err)
}

func (ts *UserUsecasesTestSuite) TestRefreshTokenNotFound() {
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOne", mock.Anything).Return(nil, errors.New("user not found"))
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("DecodeToken", mock.Anything).Return(services.JWTData{}, nil)
	ts.jwtServiceMock.On("RefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", "refresh", nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, _, err := uc.RefreshToken("sometoken")

	ts.Error(err)
	ts.IsType(&sd.InvalidRefreshToken{}, err)
}

func (ts *UserUsecasesTestSuite) TestRefreshTokenWithErrorRefreshing() {
	user := mockedUser()
	tokenData := services.JWTData{
		UserID: user.ID,
		Role:   user.Role,
	}
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	ts.userRepositoryMock.On("FindOne", mock.Anything).Return(&user, nil)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("DecodeToken", mock.Anything).Return(tokenData, nil)
	ts.jwtServiceMock.On("RefreshToken", mock.Anything, mock.Anything, mock.Anything).Return("token", "refresh", errors.New("error refreshing token"))
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	_, _, err := uc.RefreshToken("sometoken")

	ts.Error(err)
	ts.IsType(&sd.InvalidRefreshToken{}, err)
}

func (ts *UserUsecasesTestSuite) TestLoggerOut() {
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.jwtServiceMock.On("InvalidateToken", mock.Anything).Return(nil)
	uc := usecases.NewUserUseCases(
		ts.userRepositoryMock,
		ts.jwtServiceMock,
		ts.hashServiceMock,
	)

	err := uc.LoggerOut("sometoken")

	ts.NoError(err)
}

func mockedUser() domain.User {
	user := domain.NewUser()
	user.ID = "user-id"
	user.Username = "mm"
	user.FirstName = "Martín"
	user.LastName = "Martínez"
	user.Email = "mm@example.com"
	return user
}

func TestUserUsecasesTestSuite(t *testing.T) {
	ts := new(UserUsecasesTestSuite)
	ts.hashServiceMock = new(mockServices.MockHashService)
	ts.jwtServiceMock = new(mockServices.MockJWTService)
	ts.userRepositoryMock = new(mocks.MockUserRepository)
	suite.Run(t, ts)
}
