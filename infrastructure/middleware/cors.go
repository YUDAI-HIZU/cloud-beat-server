package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	allowOrigin := os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")
	if allowOrigin == "" {
		allowOrigin = "*"
	}
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, PATCH, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
