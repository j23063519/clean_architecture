package captcha

import (
	"errors"
	"time"

	"github.com/j23063519/clean_architecture/config"
	"github.com/j23063519/clean_architecture/pkg/redis"
	"github.com/spf13/cast"
)

// RedisStore accomplish base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set accomplish base64Captcha.Store interface's Set
func (rds *RedisStore) Set(key string, value string) error {
	expireTime := time.Minute * time.Duration(cast.ToInt64(config.Config.Captcha.EXPIRETIME))

	// save captcha
	if ok := rds.RedisClient.Set(rds.KeyPrefix+key, value, expireTime); !ok {
		return errors.New("saving captcha failed")
	}
	return nil
}

// Get accomplish base64Captcha.Store interface's Get
func (rds *RedisStore) Get(key string, clear bool) string {
	newKey := rds.KeyPrefix + key
	val := rds.RedisClient.Get(newKey)

	if clear {
		rds.RedisClient.Del(newKey)
	}

	return val
}

// Verify accomplish base64Captcha.Store interface's Verify
func (rds *RedisStore) Verify(key, answer string, clear bool) bool {
	v := rds.Get(key, clear)
	return v == answer
}
