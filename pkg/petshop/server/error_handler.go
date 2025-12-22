// Package server provides dtos, error handling and http router to expose petshop endpoints.
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

func handleError(c *gin.Context, err error) {
	switch err.(type) {
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	case *domain.PetNotFound:
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
}
