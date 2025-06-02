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
)
