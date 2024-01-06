package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Credentials",
			"false",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, DELETE, GET, PUT, OPTIONS",
		)
		c.Writer.Header().Set(
			"Access-Control-Expose-Headers",
			"Authorization",
		)
		c.Writer.Header().Set(
			"Access-Control-Max-Age",
			cast.ToString(12*time.Hour),
		)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
