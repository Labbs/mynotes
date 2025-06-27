package caching

import "time"

// Cache defines the interface for caching implementations
type CacheInterface interface {
	Set(key string, value interface{})
	SetWithTTL(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string) bool
	Clear()
	Close() error
}
