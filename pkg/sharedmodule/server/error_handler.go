// Package server provides shared HTTP utilities for all modules.
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
)

// ErrorResponse provides a consistent JSON error response format.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

// NewErrorResponse creates a new ErrorResponse.
func NewErrorResponse(code string, message string, details any) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// ErrorClassifier is a function that classifies an error and returns
// the HTTP status code and error code. Return (0, "") if the error is not handled.
type ErrorClassifier func(err error) (status int, code string)

// HandleErrorWithClassifier handles errors using the provided classifier function.
// If the classifier returns (0, ""), it falls back to base error handling.
func HandleErrorWithClassifier(c *gin.Context, err error, classifier ErrorClassifier) {
	// Try module-specific classifier first
	if classifier != nil {
		if status, code := classifier(err); status != 0 {
			c.JSON(status, NewErrorResponse(code, err.Error(), nil))
			return
		}
	}

	// Handle shared domain errors
	switch err.(type) {
	case *domain.InvalidRefreshToken:
		c.JSON(http.StatusUnauthorized, NewErrorResponse("INVALID_REFRESH_TOKEN", err.Error(), nil))
		return
	}

	// Default to internal server error
	c.JSON(http.StatusInternalServerError, NewErrorResponse("INTERNAL_ERROR", err.Error(), nil))
}

// SendError sends an error response with the specified status code and error code.
func SendError(c *gin.Context, status int, code string, message string) {
	c.JSON(status, NewErrorResponse(code, message, nil))
}

// SendErrorWithDetails sends an error response with additional details.
func SendErrorWithDetails(c *gin.Context, status int, code string, message string, details any) {
	c.JSON(status, NewErrorResponse(code, message, details))
}

// Common HTTP error helpers

// SendNotFound sends a 404 Not Found error response.
func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, "NOT_FOUND", message)
}

// SendUnauthorized sends a 401 Unauthorized error response.
func SendUnauthorized(c *gin.Context, message string) {
	SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", message)
}

// SendBadRequest sends a 400 Bad Request error response.
func SendBadRequest(c *gin.Context, message string) {
	SendError(c, http.StatusBadRequest, "BAD_REQUEST", message)
}

// SendInternalError sends a 500 Internal Server Error response.
func SendInternalError(c *gin.Context, message string) {
	SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", message)
}
