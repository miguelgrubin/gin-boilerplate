// Package server provides dtos, error handling and http router to expose users endpoints.
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sharedServer "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/users/domain"
)

// usersErrorClassifier classifies user-domain specific errors.
func usersErrorClassifier(err error) (int, string) {
	switch err.(type) {
	case *domain.UsernameNotFound:
		return http.StatusNotFound, "USER_NOT_FOUND"
	case *domain.InvalidLogin:
		return http.StatusUnauthorized, "INVALID_LOGIN"
	case *domain.InvalidEmail:
		return http.StatusBadRequest, "INVALID_EMAIL"
	case *domain.InvalidUsername:
		return http.StatusBadRequest, "INVALID_USERNAME"
	case *domain.InvalidPhone:
		return http.StatusBadRequest, "INVALID_PHONE"
	}
	return 0, ""
}

// handleError handles errors for the users module using the shared error handler.
func handleError(c *gin.Context, err error) {
	sharedServer.HandleErrorWithClassifier(c, err, usersErrorClassifier)
}
