package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	sDomain "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
	uMocks "github.com/miguelgrubin/gin-boilerplate/pkg/users/mocks"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createServerFixture(useCases usecases.UserUseCasesInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	os.Setenv("APP_ENV", "test")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	pc := server.NewUserHandlers(useCases)
	pc.SetupRoutes(v1)
	return router
}

func mockUser() domain.User {
	user := domain.NewUser()
	user.Username = "username"
	user.FirstName = "User Name"
	user.LastName = "Last Name"
	user.Email = "user@example.com"
	user.PasswordHash = "securepassword"
	user.Phone = "1234567890"
	user.Status = "active"
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user
}

func TestCreateUser(t *testing.T) {
	body, err := json.Marshal(server.UserCreateRequest{
		Username:  "username",
		FirstName: "User Name",
		LastName:  "Last Name",
		Email:     "user@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "username",
		FirstName: "User Name",
		LastName:  "Last Name",
		Email:     "user@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("Creator", mock.Anything).Return(user, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateUserBadRequest(t *testing.T) {
	body := []byte(`{"username": "username", "first_name": 312, "last_name": "Last Name", "email": "user@example.com", "password": "securepassword", "phone": "1234567890"}`)
	user := domain.CreateUser(domain.CreateUserParams{
		Username:  "username",
		FirstName: "User Name",
		LastName:  "Last Name",
		Email:     "user@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("Creator", mock.Anything).Return(user, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserWithErrors(t *testing.T) {
	body, err := json.Marshal(server.UserCreateRequest{
		Username:  "username",
		FirstName: "User Name",
		LastName:  "Last Name",
		Email:     "user@example.com",
		Password:  "securepassword",
		Phone:     "1234567890",
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("Creator", mock.Anything).Return(domain.User{}, errors.New("error creating user"))
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestShowUser(t *testing.T) {
	user := mockUser()
	puc := new(uMocks.MockUserUseCases)
	puc.On("Shower", mock.Anything).Return(user, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/username", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShowUserNotFound(t *testing.T) {
	puc := new(uMocks.MockUserUseCases)
	puc.On("Shower", mock.Anything).Return(domain.User{}, &domain.UsernameNotFound{Username: "username"})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/username", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateUser(t *testing.T) {
	newName := "username23"
	user := mockUser()
	user.FirstName = newName
	body, err := json.Marshal(server.UserUpdateRequest{
		FirstName: &newName,
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("Updater", mock.Anything, mock.Anything).Return(user, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/v1/user/username", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserBadRequest(t *testing.T) {
	body := []byte(`{"first_name": false}`)
	newName := "username23"
	user := mockUser()
	user.FirstName = newName
	puc := new(uMocks.MockUserUseCases)
	puc.On("Updater", mock.Anything, mock.Anything).Return(user, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/v1/user/username", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateUserNotFound(t *testing.T) {
	newName := "username23"
	body, err := json.Marshal(server.UserUpdateRequest{
		FirstName: &newName,
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("Updater", mock.Anything, mock.Anything).Return(domain.User{}, &domain.UsernameNotFound{Username: "username"})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/v1/user/username", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteUser(t *testing.T) {
	puc := new(uMocks.MockUserUseCases)
	puc.On("Deleter", mock.Anything).Return(nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/user/username", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteUserNotFound(t *testing.T) {
	puc := new(uMocks.MockUserUseCases)
	puc.On("Deleter", mock.Anything).Return(&domain.UsernameNotFound{Username: "username"})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/user/username", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestLoginUser(t *testing.T) {
	body, err := json.Marshal(server.UserLoginRequest{
		Username: "username",
		Password: "securepassword",
	})
	token := "jwt.token.here"
	refreshToken := "refresh.token.here"
	puc := new(uMocks.MockUserUseCases)
	puc.On("LoggerIn", mock.Anything, mock.Anything).Return(token, refreshToken, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUserBadRequest(t *testing.T) {
	body := []byte(`{"ole": false, "password": "securepassword"}`)
	token := "jwt.token.here"
	refreshToken := "refresh.token.here"
	puc := new(uMocks.MockUserUseCases)
	puc.On("LoggerIn", mock.Anything, mock.Anything).Return(token, refreshToken, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginUserInvalid(t *testing.T) {
	body, err := json.Marshal(server.UserLoginRequest{
		Username: "username",
		Password: "securepassword",
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("LoggerIn", mock.Anything, mock.Anything).Return("", "", &domain.InvalidLogin{})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRefreshToken(t *testing.T) {
	body, err := json.Marshal(server.UserRefreshTokenRequest{
		RefreshToken: "old.refresh.token.here",
	})
	token := "jwt.token.here"
	refreshToken := "refresh.token.here"
	puc := new(uMocks.MockUserUseCases)
	puc.On("RefreshToken", mock.Anything).Return(token, refreshToken, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/refresh", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRefreshTokenBadRequest(t *testing.T) {
	body := []byte(`{"refresh_token": false}`)
	token := "jwt.token.here"
	refreshToken := "refresh.token.here"
	puc := new(uMocks.MockUserUseCases)
	puc.On("RefreshToken", mock.Anything).Return(token, refreshToken, nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/refresh", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRefreshTokenInvalid(t *testing.T) {
	body, err := json.Marshal(server.UserRefreshTokenRequest{
		RefreshToken: "old.refresh.token.here",
	})
	puc := new(uMocks.MockUserUseCases)
	puc.On("RefreshToken", mock.Anything).Return("", "", &sDomain.InvalidRefreshToken{})
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/refresh", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLogout(t *testing.T) {
	puc := new(uMocks.MockUserUseCases)
	puc.On("LoggerOut", mock.Anything).Return(nil)
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	ctx := context.WithValue(context.Background(), "Authorization", "Bearer jwt.token.here")
	req, _ := http.NewRequestWithContext(ctx, "POST", "/v1/auth/logout", bytes.NewBuffer([]byte{}))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestLogoutWithError(t *testing.T) {
	puc := new(uMocks.MockUserUseCases)
	puc.On("LoggerOut", mock.Anything).Return(errors.New("error logging out"))
	router := createServerFixture(puc)
	w := httptest.NewRecorder()
	ctx := context.WithValue(context.Background(), "Authorization", "Bearer jwt.token.here")
	req, _ := http.NewRequestWithContext(ctx, "POST", "/v1/auth/logout", bytes.NewBuffer([]byte{}))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
