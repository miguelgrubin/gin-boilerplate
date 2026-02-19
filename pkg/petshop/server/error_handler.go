// Package server provides dtos, error handling and http router to expose petshop endpoints.
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	sharedServer "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/server"
)

// petshopErrorClassifier classifies petshop-domain specific errors.
func petshopErrorClassifier(err error) (int, string) {
	switch err.(type) {
	case *domain.PetNotFound:
		return http.StatusNotFound, "PET_NOT_FOUND"
	}
	return 0, ""
}

// handleError handles errors for the petshop module using the shared error handler.
func handleError(c *gin.Context, err error) {
	sharedServer.HandleErrorWithClassifier(c, err, petshopErrorClassifier)
}
