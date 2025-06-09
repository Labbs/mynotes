package caching

import (
	"fmt"
	"time"

	"github.com/labbs/mynotes/pkg/config"
	"github.com/rs/zerolog"
)

var Cache CacheInterface

type Config struct {
	Logger      zerolog.Logger
	CacheConfig CacheConfig
}

type CacheType string

const (
	MemoryCacheType CacheType = "memory"
	RedisCacheType  CacheType = "redis"
)

type CacheConfig struct {
	Type       CacheType
	MaxSize    int // Pour le cache mémoire uniquement
	DefaultTTL time.Duration

	// Configuration Redis
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func (c *Config) Configure() error {
	var err error
	// Configure caching settings here if needed
	c.Logger.Info().Msg("Caching configuration initialized")
	switch config.Cache.Type {
	case string(MemoryCacheType):
		c.Logger.Info().Msgf("Using Memory Cache with MaxSize: %d and DefaultTTL: %d", config.Cache.Memory.Size, config.Cache.Expire)
		Cache = NewMemoryCache(config.Cache.Memory.Size, time.Duration(config.Cache.Expire))
	case string(RedisCacheType):
		c.Logger.Info().Msgf("Using Redis Cache at %s with DB: %d and DefaultTTL: %d", config.Cache.Redis.Addr, config.Cache.Redis.DB, config.Cache.Expire)
		Cache = NewRedisCache(config.Cache.Redis.Addr, config.Cache.Redis.Password, config.Cache.Redis.DB, time.Duration(config.Cache.Expire))
	default:
		c.Logger.Error().Msgf("Unsupported cache type: %s", config.Cache.Type)
		err = fmt.Errorf("unsupported cache type: %s", config.Cache.Type)
	}
	return err
}
