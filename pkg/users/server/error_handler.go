// Package server provides dtos, error handling and http router to expose petshop endpoints.
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sDomain "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
)

func handleError(c *gin.Context, err error) {
	switch err.(type) {
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	case *domain.UsernameNotFound:
		c.JSON(http.StatusNotFound, err.Error())
		return
	case *domain.InvalidLogin:
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	case *sDomain.InvalidRefreshToken:
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
}
