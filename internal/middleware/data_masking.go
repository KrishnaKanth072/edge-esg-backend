package middleware

import (
	"github.com/gin-gonic/gin"
)

func DataMasking() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("user_roles")
		if !exists {
			c.Set("user_role", "TRADER")
		} else {
			userRoles := roles.([]string)
			if len(userRoles) > 0 {
				c.Set("user_role", userRoles[0])
			} else {
				c.Set("user_role", "TRADER")
			}
		}
		c.Next()
	}
}
