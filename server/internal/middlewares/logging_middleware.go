package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/bermanbenjamin/futStats/internal/logger"
)

// RequestIDKey is the key used to store request ID in context
const RequestIDKey = "request_id"

// LoggingMiddleware creates a middleware for HTTP request/response logging
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID
		requestID := uuid.New().String()
		c.Set(RequestIDKey, requestID)

		// Add request ID to response headers
		c.Header("X-Request-ID", requestID)

		// Start timer
		start := time.Now()

		// Get request details
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()
		remoteAddr := c.ClientIP()

		// Create logger with request context
		reqLogger := logger.GetGlobal().WithFields(
			zap.String("request_id", requestID),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("user_agent", userAgent),
			zap.String("remote_addr", remoteAddr),
		)

		// Log request start
		reqLogger.Info("Request started")

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Get response details
		statusCode := c.Writer.Status()

		// Log request completion
		reqLogger.LogHTTPRequest(method, path, userAgent, remoteAddr, statusCode, duration, requestID)

		// Log errors if any
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				reqLogger.Error("Request error", zap.Error(err))
			}
		}
	}
}

// GetRequestID extracts request ID from context
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}

// GetRequestLogger returns a logger with request context
func GetRequestLogger(c *gin.Context) *logger.Logger {
	requestID := GetRequestID(c)
	if requestID != "" {
		return logger.GetGlobal().WithRequestID(requestID)
	}
	return logger.GetGlobal()
}
