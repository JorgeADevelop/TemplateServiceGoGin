package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" || c.Request.Method == "DELETE" {
			csrfToken := c.Request.Header.Get("X-CSRF-Token")
			expectedToken, err := c.Cookie("csrf_token")
			if err != nil || csrfToken != expectedToken {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CSRF token inválido"})
				return
			}
		}
		c.Next()
	}
}
