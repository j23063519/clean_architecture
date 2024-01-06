package init

import (
	"fmt"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/cache"
)

func SetCache() {
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v",
			config.Config.Redis.HOST, config.Config.Redis.PORT,
		),
		"",
		"",
		config.Config.Redis.CACHEDB,
	)
	cache.InitWithCacheStore(rds)
}
