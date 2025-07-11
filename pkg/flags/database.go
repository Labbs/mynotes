package flags

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func DatabaseFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.dsn",
			Aliases:     []string{"dsn"},
			EnvVars:     []string{"DATABASE_DSN"},
			Usage:       "Database Gorm DSN",
			Value:       "./database.db",
			Destination: &config.Database.DSN,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.dialect",
			Aliases:     []string{"dialect"},
			EnvVars:     []string{"DATABASE_DIALECT"},
			Usage:       "Database Gorm Dialect",
			Value:       "sqlite",
			Destination: &config.Database.Dialect,
		}),
	}
}
