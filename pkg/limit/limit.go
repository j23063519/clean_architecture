package limit

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/j23063519/clean_architecture/pkg/redis"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

// transform url "/" to "-"
// transform url ":" to "-"
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}

func CheckRate(c *gin.Context, key, formatted string) (limiterlib.Context, error) {
	// new limiter then get limiter.Rate
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		log.ErrorJSON("limit rate[limiter]", "CheckRate", err)
		return context, err
	}

	// init storage and use redis.Redis
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		// set prefix to identifier
		Prefix: config.Config.App.NAME + ":limiter",
	})
	if err != nil {
		log.ErrorJSON("limit rate[limiter]", "CheckRate", err)
		return context, err
	}

	// new limiter
	limiterObj := limiterlib.New(store, rate)

	// get result from limiter-once
	if c.GetBool("limiter-once") {
		// get result but not increase request time
		return limiterObj.Peek(c, key)
	} else {

		// when use LimitIP in multiple route only increase one request time
		c.Set("limiter-once", true)

		// get result and increment request time
		return limiterObj.Get(c, key)
	}
}
