package flags

import (
	"github.com/labbs/mynotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func LoggerFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "logger.level",
			Aliases:     []string{"logl"},
			EnvVars:     []string{"LOG_LEVEL"},
			Usage:       "Set the logging level (debug, info, warn, error, fatal, panic)",
			Value:       "info",
			Destination: &config.Logger.Level,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "logger.pretty",
			Aliases:     []string{"logp"},
			EnvVars:     []string{"LOG_PRETTY"},
			Usage:       "Enable pretty logging",
			Value:       false,
			Destination: &config.Logger.Pretty,
		}),
	}
}
