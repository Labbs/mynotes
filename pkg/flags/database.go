package flags

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func DatabaseFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.gorm.dsn",
			Aliases:     []string{"dgd"},
			EnvVars:     []string{"DATABASE_GORM_DSN"},
			Usage:       "Database Gorm DSN",
			Value:       "./database.db",
			Destination: &config.Database.DSN,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.gorm.dialect",
			Aliases:     []string{"dgdialect"},
			EnvVars:     []string{"DATABASE_GORM_DIALECT"},
			Usage:       "Database Gorm Dialect",
			Value:       "sqlite",
			Destination: &config.Database.Dialect,
		}),
	}
}
