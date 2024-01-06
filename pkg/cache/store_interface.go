package cache

import "time"

type Store interface {
	Set(key, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key, value string)
	Flush()

	IsAlive() error

	// if parameter only one key's value increase one
	//
	// if parameter have two and first is key and second is value which will increase value
	Increment(parameters ...interface{})

	// if parameter only one key's value decrease one
	//
	// if parameter have two and first is key and second is value which will decrease value
	Decrement(parameters ...interface{})
}
