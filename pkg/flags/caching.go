package flags

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func CachingFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "caching.enabled",
			Aliases:     []string{"ce"},
			EnvVars:     []string{"CACHING_ENABLED"},
			Usage:       "Enable caching",
			Value:       true,
			Destination: &config.Cache.Enable,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "caching.type",
			Aliases:     []string{"ct"},
			EnvVars:     []string{"CACHING_TYPE"},
			Usage:       "Type of caching to use (e.g., 'memory', 'redis')",
			Value:       "memory",
			Destination: &config.Cache.Type,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "caching.expire",
			Aliases:     []string{"cexp"},
			EnvVars:     []string{"CACHING_EXPIRE"},
			Usage:       "Cache expiration time in seconds",
			Value:       60 * 10, // Default to 10 minutes
			Destination: &config.Cache.Expire,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "caching.redis.addr",
			Aliases:     []string{"cra"},
			EnvVars:     []string{"CACHING_REDIS_ADDR"},
			Usage:       "Redis server address",
			Value:       "localhost:6379",
			Destination: &config.Cache.Redis.Addr,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "caching.redis.password",
			Aliases:     []string{"crp"},
			EnvVars:     []string{"CACHING_REDIS_PASSWORD"},
			Usage:       "Redis server password",
			Value:       "",
			Destination: &config.Cache.Redis.Password,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "caching.redis.db",
			Aliases:     []string{"crdb"},
			EnvVars:     []string{"CACHING_REDIS_DB"},
			Usage:       "Redis database number",
			Value:       0,
			Destination: &config.Cache.Redis.DB,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "caching.memory.size",
			Aliases:     []string{"cms"},
			EnvVars:     []string{"CACHING_MEMORY_SIZE"},
			Usage:       "Memory cache size in MB",
			Value:       100, // Default to 100 MB
			Destination: &config.Cache.Memory.Size,
		}),
	}
}
