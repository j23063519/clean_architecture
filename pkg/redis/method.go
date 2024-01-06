package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/j23063519/clean_architecture/pkg/console"
	"github.com/j23063519/clean_architecture/pkg/log"
)

// Ping 用以測試 redis 連接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 儲存 key 對應的 value，且設置 過期時間
func (rds RedisClient) Set(key string, val interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, val, expiration).Err(); err != nil {
		log.ErrorJSON("Redis", "Set", err.Error())
		console.Error("Redis[Set], error:" + err.Error())
		return false
	}
	return true
}

// Get 獲取 key 對應的值
func (rds RedisClient) Get(key string) string {
	res, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.ErrorJSON("Redis", "Get", err.Error())
			console.Error("Redis[Get], error:" + err.Error())
		}
		return ""
	}
	return res
}

// Has 判斷 key 是否存在，內部錯誤 和 redis.Nil 都回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.ErrorJSON("RedisClient", "Has", err)
		}
		return false
	}
	return true
}

// Del 刪除存儲在 redis 裡的數據，支持多個 key 參數
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		log.ErrorJSON("Redis", "Del", err.Error())
		console.Error("Redis[Del], error:" + err.Error())
		return false
	}
	return true
}

// FlushDB 清空當前 redis db 的所有數據
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		log.ErrorJSON("Redis", "FlushDB", err.Error())
		console.Error("Redis[FlushDB], error:" + err.Error())
		return false
	}
	return true
}

// Increment 當參數只有1個時，為 key，其值加 1
//
// 當參數有 2 個時，第一個參數為 key，第二個參數為要增加的值 int64 類型
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			log.ErrorJSON("Redis", "Increment", err.Error())
			console.Error("Redis[Increment], error:" + err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		val := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, val).Err(); err != nil {
			log.ErrorJSON("Redis", "Increment", err.Error())
			console.Error("Redis[Increment], error:" + err.Error())
			return false
		}
	default:
		log.ErrorJSON("Redis", "Increment", "參數過多")
		console.Error("Redis[Increment], error:" + "參數過多")
		return false
	}
	return true
}

// Decrement 當參數只有1個時，為 key，其值減 1
//
// 當參數有 2 個時，第一個參數為 key，第二個參數為要減去的值 int64 類型
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			log.ErrorJSON("Redis", "Decrement", err.Error())
			console.Error("Redis[Decrement], error:" + err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		val := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, val).Err(); err != nil {
			log.ErrorJSON("Redis", "Decrement", err.Error())
			console.Error("Redis[Decrement], error:" + err.Error())
			return false
		}
	default:
		log.ErrorJSON("Redis", "Decrement", "參數過多")
		console.Error("Redis[Decrement], error:" + "參數過多")
		return false
	}
	return true
}
