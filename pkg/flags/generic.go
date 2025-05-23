package flags

import (
	"github.com/labbs/mynotes/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func GenericFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"CONFIG"},
			Usage:   "Load configuration from `FILE`",
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "dev-mode",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DEV_MODE"},
			Usage:       "Enable development mode",
			Value:       false,
			Destination: &config.DevMode,
		}),
	}
}
