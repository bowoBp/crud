package midleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware untuk otentikasi dengan Bearer token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer your_token_here" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
