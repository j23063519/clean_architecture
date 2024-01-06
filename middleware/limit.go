package middleware

import (
	"github.com/gin-gonic/gin"
	limitpkg "github.com/j23063519/clean_architecture/pkg/limit"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/j23063519/clean_architecture/pkg/response"
	"github.com/spf13/cast"
)

// LimitIP global limit middleware
//
// * 5 reqs/second: "5-S"
//
// * 10 reqs/minute: "10-M"
//
// * 1000 reqs/hour: "1000-H"
//
// * 2000 reqs/day: "2000-D"
func LimitIP(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ip rate limit
		ip := limitpkg.GetIP(c)
		if ok := limitHandler(c, ip, limit); !ok {
			return
		}
		c.Next()
	}
}

// global limit per ip with route middleware
func LimitPerRoute(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("limiter-once", false)

		key := limitpkg.GetRouteWithIP(c)
		if !limitHandler(c, key, limit) {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	// get rate
	rate, err := limitpkg.CheckRate(c, key, limit)

	if err != nil {
		log.ErrorJSON("limitHandler", "limit", err)
		response.Response(c, 500, "limit: "+err.Error(), nil)
		return false
	}

	// setting header information
	// X-RateLimit-Limit: 10000
	// X-RateLimit-Remaining: 9993
	// X-RateLimit-Reset: 1513784506: if current timestamp is greater than 1513784506 then X-RateLimit-Remaining will be reset
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached {
		response.Response(c, 429, "", nil)
		return false
	}

	return true
}
