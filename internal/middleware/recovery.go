package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indraplrg/technical_test/internal/response"
)

// Recovery catches panics, logs them and returns a consistent JSON error.
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		slog.Error("panic recovered",
			"error", recovered,
			"request_id", c.GetString("request_id"),
			"uri", c.Request.RequestURI,
		)
		response.Error(c, http.StatusInternalServerError, "internal server error")
		c.Abort()
	})
}
