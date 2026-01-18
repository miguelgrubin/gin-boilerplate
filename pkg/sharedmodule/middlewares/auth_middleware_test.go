package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/middlewares"
	sMocks "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createServerFixture() *gin.Engine {
	gin.SetMode(gin.TestMode)
	os.Setenv("APP_ENV", "test")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router
}

func TestAuthMiddlewareAccept(t *testing.T) {
	router := createServerFixture()
	js := new(sMocks.MockJWTService)
	js.On("ValidateToken", mock.Anything).Return(true)
	router.Use(middlewares.AuthRequired(js))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "access granted"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Add("Authorization", "Bearer jwt.token.here")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "access granted")

}

func TestAuthMiddlewareDeny(t *testing.T) {
	router := createServerFixture()
	js := new(sMocks.MockJWTService)
	js.On("ValidateToken", mock.Anything).Return(false)
	router.Use(middlewares.AuthRequired(js))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "access granted"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Add("Authorization", "Bearer jwt.token.here")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Contains(t, w.Body.String(), "unauthorized")
}

func TestAuthMiddlewareWithoutToken(t *testing.T) {
	router := createServerFixture()
	js := new(sMocks.MockJWTService)
	js.On("ValidateToken", mock.Anything).Return(false)
	router.Use(middlewares.AuthRequired(js))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "access granted"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Add("Authorization", "Bearer")
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Contains(t, w.Body.String(), "unauthorized")
}
