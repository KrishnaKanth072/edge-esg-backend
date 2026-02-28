package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/edgeesg/edge-esg-backend/internal/error_codes"
)

type KeycloakMiddleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewKeycloakMiddleware(issuerURL, clientID string) (*KeycloakMiddleware, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
	return &KeycloakMiddleware{verifier: verifier}, nil
}

func (k *KeycloakMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    error_codes.AuthUnauthorized,
				"message": "Missing authorization header",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    error_codes.AuthInvalidToken,
				"message": "Invalid token format",
			})
			c.Abort()
			return
		}

		idToken, err := k.verifier.Verify(context.Background(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    error_codes.AuthExpiredToken,
				"message": "Token verification failed",
			})
			c.Abort()
			return
		}

		var claims struct {
			Email string   `json:"email"`
			Roles []string `json:"roles"`
		}
		if err := idToken.Claims(&claims); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    error_codes.AuthInvalidToken,
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		c.Set("user_email", claims.Email)
		c.Set("user_roles", claims.Roles)
		c.Next()
	}
}

func (k *KeycloakMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("user_roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    error_codes.AuthForbidden,
				"message": "No roles found",
			})
			c.Abort()
			return
		}

		userRoles := roles.([]string)
		hasRole := false
		for _, r := range userRoles {
			if r == role || r == "ADMIN" {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    error_codes.AuthForbidden,
				"message": "Insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
