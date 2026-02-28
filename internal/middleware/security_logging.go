package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// SecurityLogger logs security-relevant events
func SecurityLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Process request
		c.Next()

		// Log after request
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Get user info if authenticated
		userEmail, _ := c.Get("user_email")
		userRoles, _ := c.Get("user_roles")

		logEvent := log.Info().
			Str("method", method).
			Str("path", path).
			Int("status", statusCode).
			Dur("latency", latency).
			Str("client_ip", clientIP).
			Str("user_agent", userAgent)

		if userEmail != nil {
			logEvent = logEvent.Str("user_email", userEmail.(string))
		}

		if userRoles != nil {
			logEvent = logEvent.Interface("user_roles", userRoles)
		}

		// Log security events
		if statusCode == 401 {
			logEvent.Msg("SECURITY: Unauthorized access attempt")
		} else if statusCode == 403 {
			logEvent.Msg("SECURITY: Forbidden access attempt")
		} else if statusCode >= 400 && statusCode < 500 {
			logEvent.Msg("SECURITY: Client error")
		} else if statusCode >= 500 {
			logEvent.Msg("SECURITY: Server error")
		} else {
			logEvent.Msg("Request processed")
		}
	}
}

// AuditLog logs critical security events
func AuditLog(event string, details map[string]interface{}) {
	logEvent := log.Info().
		Str("event_type", "AUDIT").
		Str("event", event).
		Time("timestamp", time.Now())

	for key, value := range details {
		logEvent = logEvent.Interface(key, value)
	}

	logEvent.Msg("Audit event")
}

// LogAuthenticationAttempt logs authentication attempts
func LogAuthenticationAttempt(success bool, email string, ip string, reason string) {
	if success {
		log.Info().
			Str("event_type", "AUTH_SUCCESS").
			Str("email", email).
			Str("ip", ip).
			Msg("Successful authentication")
	} else {
		log.Warn().
			Str("event_type", "AUTH_FAILURE").
			Str("email", email).
			Str("ip", ip).
			Str("reason", reason).
			Msg("Failed authentication attempt")
	}
}

// LogDataAccess logs access to sensitive data
func LogDataAccess(userEmail string, dataType string, bankID string, action string) {
	log.Info().
		Str("event_type", "DATA_ACCESS").
		Str("user_email", userEmail).
		Str("data_type", dataType).
		Str("bank_id", bankID).
		Str("action", action).
		Time("timestamp", time.Now()).
		Msg("Sensitive data access")
}
