package redis

import (
	"context"
	"sync"

	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/redis/go-redis/v9"
)

// RedisClient service
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// only new one time (singleton pattern)
var once sync.Once

// global Redis， DB:1
var Redis *RedisClient

// connet Redis and new redis client
func ConnectRedis(addr, user, pwd string, db int) {
	once.Do(func() {
		Redis = NewClient(addr, user, pwd, db)
	})
}

// new redis client
func NewClient(addr, user, pwd string, db int) (rds *RedisClient) {
	// init RedisClient
	rds = &RedisClient{}
	// use default context
	rds.Context = context.Background()

	// use redis library init connection
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: pwd,
		DB:       db,
	})

	// test connection
	if err := rds.Ping(); err != nil {
		console.Error("Redis connection error: " + err.Error())
		log.ErrorJSON("Redis", "REDIS 連接", err)
	}

	return
}
