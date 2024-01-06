package init

import (
	"net/http"
	"strings"

	"github.com/j23063519/clean_architecture/middleware"
	"github.com/j23063519/clean_architecture/pkg/response"
	"github.com/j23063519/clean_architecture/route"

	"github.com/gin-gonic/gin"
)

// set route
func SetRoute(router *gin.Engine) {
	// register global middleware
	registerGlobalMiddleware(router)

	// register handler
	route.RegisterHandler(router)

	// set 404 route
	setup404Handler(router)
}

// register common middleware
func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middleware.CORS(),
		// middleware.Log(),
		// middleware.Recovery(),
	)
}

// 404 route not found
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// get headers from request
		acceptString := c.Request.Header.Get("Accept")
		// if it is html
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "Go Back 404")
		} else {
			response.Response(c, 404, "route not found", nil)
		}
	})
}
