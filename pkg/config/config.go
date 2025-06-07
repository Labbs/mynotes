package config

var (
	DevMode bool

	Database struct {
		Dialect string // Database engine (sqlite, postgres, etc.)
		DSN     string
	}

	Logger struct {
		Level  string
		Pretty bool
	}

	Server struct {
		Port     int
		HttpLogs bool
	}

	Session struct {
		SecretKey string
		Expire    int
		Issuer    string
	}

	Document struct {
		ExcalidrawLibsPath string
	}

	Cache struct {
		Type   string // Cache type (memory, redis, etc.)
		Expire int    // Cache expiration time in seconds
		Redis  struct {
			Addr     string // Redis server address
			Password string // Redis password
			DB       int    // Redis database number
		}

		Memory struct {
			Size int // Memory cache size in MB
		}

		Enable bool // Enable or disable caching
	}
)
