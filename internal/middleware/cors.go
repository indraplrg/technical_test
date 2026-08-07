package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS returns a middleware that applies permissive CORS headers.
// When allowedOrigin is empty, "*" is used.
func CORS(allowedOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := allowedOrigin
		if origin == "" {
			origin = "*"
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-ID")
		c.Header("Access-Control-Expose-Headers", "X-Request-ID")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
