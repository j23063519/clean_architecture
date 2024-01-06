package cache

import (
	"time"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/redis"
)

// cache.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(addr, user, pwd string, db int) *RedisStore {
	// init RedisClient
	rds := &RedisStore{}
	rds.RedisClient = redis.NewClient(addr, user, pwd, db)
	rds.KeyPrefix = config.Config.App.NAME + ":cache:"

	return rds
}

func (rds *RedisStore) Set(key, value string, expireTime time.Duration) {
	rds.RedisClient.Set(rds.KeyPrefix+key, value, expireTime)
}

func (rds *RedisStore) Get(key string) string {
	return rds.RedisClient.Get(rds.KeyPrefix + key)
}

func (rds *RedisStore) Has(key string) bool {
	return rds.RedisClient.Has(rds.KeyPrefix + key)
}

func (rds *RedisStore) Forget(key string) {
	rds.RedisClient.Del(rds.KeyPrefix + key)
}

func (rds *RedisStore) Forever(key, value string) {
	rds.RedisClient.Set(rds.KeyPrefix+key, value, 0)
}

func (rds *RedisStore) Flush() {
	rds.RedisClient.FlushDB()
}

func (rds *RedisStore) IsAlive() error {
	return rds.RedisClient.Ping()
}

func (rds *RedisStore) Increment(parameters ...interface{}) {
	rds.RedisClient.Increment(parameters...)
}

func (rds *RedisStore) Decrement(parameters ...interface{}) {
	rds.RedisClient.Decrement(parameters...)
}
