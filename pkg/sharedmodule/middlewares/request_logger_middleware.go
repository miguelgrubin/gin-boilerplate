package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
)

// CorrelationIDHeader is the header name for the correlation ID.
const CorrelationIDHeader = "X-Correlation-ID"

// CorrelationIDKey is the context key for storing the correlation ID.
const CorrelationIDKey = "correlation_id"

// RequestLogger returns a middleware that logs HTTP requests with structured logging.
// It adds a correlation ID to each request for tracing purposes.
func RequestLogger(logger services.LoggerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get or generate correlation ID
		correlationID := c.GetHeader(CorrelationIDHeader)
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		// Store correlation ID in context and response header
		c.Set(CorrelationIDKey, correlationID)
		c.Header(CorrelationIDHeader, correlationID)

		// Create request-scoped logger with correlation ID
		requestLogger := logger.With(
			services.String("correlation_id", correlationID),
			services.String("method", c.Request.Method),
			services.String("path", c.Request.URL.Path),
		)

		// Record start time
		start := time.Now()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log request completion
		status := c.Writer.Status()
		fields := []services.Field{
			services.Int("status", status),
			services.Duration("duration", duration),
			services.String("client_ip", c.ClientIP()),
		}

		// Add error if present
		if len(c.Errors) > 0 {
			fields = append(fields, services.String("errors", c.Errors.String()))
		}

		// Log based on status code
		switch {
		case status >= 500:
			requestLogger.Error("request completed with server error", fields...)
		case status >= 400:
			requestLogger.Warn("request completed with client error", fields...)
		default:
			requestLogger.Info("request completed", fields...)
		}
	}
}

// GetCorrelationID retrieves the correlation ID from the Gin context.
func GetCorrelationID(c *gin.Context) string {
	if id, exists := c.Get(CorrelationIDKey); exists {
		if correlationID, ok := id.(string); ok {
			return correlationID
		}
	}
	return ""
}
