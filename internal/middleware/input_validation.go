package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/edgeesg/edge-esg-backend/internal/error_codes"
	"github.com/gin-gonic/gin"
)

var (
	// SQL injection patterns
	sqlInjectionPattern = regexp.MustCompile(`(?i)(union|select|insert|update|delete|drop|create|alter|exec|execute|script|javascript|<script)`)

	// XSS patterns
	xssPattern = regexp.MustCompile(`(?i)(<script|javascript:|onerror=|onload=|<iframe|<object|<embed)`)

	// Path traversal patterns
	pathTraversalPattern = regexp.MustCompile(`(\.\./|\.\.\\|%2e%2e)`)
)

// InputValidation validates and sanitizes all request inputs
func InputValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate query parameters
		for key, values := range c.Request.URL.Query() {
			for _, value := range values {
				if !isValidInput(value) {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    error_codes.ValidationFailed,
						"message": "Invalid input detected in query parameter: " + key,
					})
					c.Abort()
					return
				}
			}
		}

		// Validate headers (except standard ones)
		dangerousHeaders := []string{"X-Custom", "X-Api-Key"}
		for _, header := range dangerousHeaders {
			if value := c.GetHeader(header); value != "" {
				if !isValidInput(value) {
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    error_codes.ValidationFailed,
						"message": "Invalid input detected in header: " + header,
					})
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

// isValidInput checks for malicious patterns
func isValidInput(input string) bool {
	// Check for SQL injection
	if sqlInjectionPattern.MatchString(input) {
		return false
	}

	// Check for XSS
	if xssPattern.MatchString(input) {
		return false
	}

	// Check for path traversal
	if pathTraversalPattern.MatchString(input) {
		return false
	}

	return true
}

// SanitizeString removes potentially dangerous characters
func SanitizeString(input string) string {
	// Remove null bytes
	input = strings.ReplaceAll(input, "\x00", "")

	// Trim whitespace
	input = strings.TrimSpace(input)

	return input
}

// RequestSizeLimit limits request body size to prevent DoS
func RequestSizeLimit(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}
