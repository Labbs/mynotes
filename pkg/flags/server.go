package flags

import (
	"github.com/labbs/mynotes/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func ServerFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "http.port",
			Aliases:     []string{"hp"},
			EnvVars:     []string{"HTTP_PORT"},
			Usage:       "Port to listen on for HTTP requests",
			Value:       8080,
			Destination: &config.Server.Port,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "http.http_logs",
			Aliases:     []string{"hl"},
			EnvVars:     []string{"HTTP_HTTP_LOGS"},
			Usage:       "Enable HTTP request logging",
			Value:       false,
			Destination: &config.Server.HttpLogs,
		}),
	}
}
