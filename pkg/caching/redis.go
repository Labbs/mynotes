package caching

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client     *redis.Client
	defaultTTL time.Duration
	ctx        context.Context
}

func NewRedisCache(addr, password string, db int, defaultTTL time.Duration) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client:     rdb,
		defaultTTL: defaultTTL,
		ctx:        context.Background(),
	}
}

func (c *RedisCache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, c.defaultTTL)
}

func (c *RedisCache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	data, err := json.Marshal(value)
	if err != nil {
		return // Or handle error appropriately
	}

	c.client.Set(c.ctx, key, data, ttl)
}

func (c *RedisCache) Get(key string) (interface{}, bool) {
	val, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return nil, false
	}

	var result interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, false
	}

	return result, true
}

func (c *RedisCache) Delete(key string) bool {
	result := c.client.Del(c.ctx, key)
	return result.Val() > 0
}

func (c *RedisCache) Clear() {
	c.client.FlushDB(c.ctx)
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}
