package cache

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/spf13/cast"
)

type CacheService struct {
	Store Store
}

// only new one time (singleton pattern)
var once sync.Once

var Cache *CacheService

// init Cache
func InitWithCacheStore(store Store) {
	once.Do(func() {
		Cache = &CacheService{
			Store: store,
		}
	})
}

func Set(key string, obj interface{}, expireTime time.Duration) {
	jsonData, err := json.Marshal(&obj)
	if err != nil {
		log.ErrorJSON("Cache", "Set", err)
	}
	Cache.Store.Set(key, string(jsonData), expireTime)
}

func Get(key string) interface{} {
	stringValue := Cache.Store.Get(key)
	var wanted interface{}
	if err := json.Unmarshal([]byte(stringValue), &wanted); err != nil {
		log.ErrorJSON("Cache", "Get", err)
	}
	return wanted
}

func Has(key string) bool {
	return Cache.Store.Has(key)
}

// model := domain.Admin{}
// cache.GetObj("key", &model)
func GetObj(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		if err := json.Unmarshal([]byte(val), &wanted); err != nil {
			log.ErrorJSON("Cache", "GetObj", err)
		}
	}
}

func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

func GetFloat64(key string) float64 {
	return cast.ToFloat64(Get(key))
}

func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}

func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}

func Forget(key string) {
	Cache.Store.Forget(key)
}

func Forever(key string, value string) {
	Cache.Store.Set(key, value, 0)
}

func Flush() {
	Cache.Store.Flush()
}

func Increment(parameters ...interface{}) {
	Cache.Store.Increment(parameters...)
}

func Decrement(parameters ...interface{}) {
	Cache.Store.Decrement(parameters...)
}

func IsAlive() error {
	return Cache.Store.IsAlive()
}
