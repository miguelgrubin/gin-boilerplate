package pkg_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/stretchr/testify/assert"
)

func createServerFixture(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)
	os.Setenv("APP_ENV", "test")
	router := pkg.SetupRouter()
	return router
}

func TestHealthcheck(t *testing.T) {
	router := createServerFixture(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Health check!", w.Body.String())
}
