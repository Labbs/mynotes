package caching

import (
	"sync"
	"time"
)

type CacheEntry struct {
	value      interface{}
	expiration time.Time
}

type MemoryCache struct {
	mutex      sync.RWMutex
	data       map[string]CacheEntry
	maxSize    int
	defaultTTL time.Duration
}

func NewMemoryCache(maxSize int, defaultTTL time.Duration) *MemoryCache {
	cache := &MemoryCache{
		data:       make(map[string]CacheEntry),
		maxSize:    maxSize,
		defaultTTL: defaultTTL,
	}

	// Start the cleanup loop in a separate goroutine
	go cache.cleanupLoop()

	return cache
}

func (c *MemoryCache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, c.defaultTTL)
}

func (c *MemoryCache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// If the cache is full and the key does not exist, remove the oldest entry
	if len(c.data) >= c.maxSize && c.data[key] == (CacheEntry{}) {
		c.removeOldest()
	}

	c.data[key] = CacheEntry{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}

	// Check if the entry has expired
	if time.Now().After(entry.expiration) {
		return nil, false
	}

	return entry.value, true
}

func (c *MemoryCache) Delete(key string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.data[key]; exists {
		delete(c.data, key)
		return true
	}
	return false
}

func (c *MemoryCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[string]CacheEntry)
}

func (c *MemoryCache) Close() error {
	// Nothing specific to close for memory cache
	return nil
}

func (c *MemoryCache) removeOldest() {
	var oldestKey string
	var oldestTime time.Time

	// Find the oldest entry
	for k, v := range c.data {
		if oldestTime.IsZero() || v.expiration.Before(oldestTime) {
			oldestKey = k
			oldestTime = v.expiration
		}
	}

	if oldestKey != "" {
		delete(c.data, oldestKey)
	}
}

func (c *MemoryCache) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *MemoryCache) cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.After(entry.expiration) {
			delete(c.data, key)
		}
	}
}
