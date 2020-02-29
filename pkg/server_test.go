package pkg

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {

	Convey("Given a server instance", t, func() {
		os.Setenv("APP_ENV", "test")
		router := setupRouter()

		Convey("it has healtcheck endpoint", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/health", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
			assert.Equal(t, "Health check", w.Body.String())
		})

	})

}
