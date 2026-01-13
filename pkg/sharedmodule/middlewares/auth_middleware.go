package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
)

func AuthRequired(js services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
		reqToken := strings.Split(bearerToken, " ")[1]
		isValid := js.ValidateToken(reqToken)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}
	}
}
