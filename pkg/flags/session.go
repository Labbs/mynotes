package flags

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// SessionFlags returns a slice of cli.Flag for session configuration.
// It's used to set up session-related flags for the CLI application.
func SessionFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.secret-key",
			Aliases:     []string{"ssk"},
			EnvVars:     []string{"SESSION_SECRET_KEY"},
			Usage:       "Session secret key",
			Value:       "zotion-secret-key",
			Destination: &config.Session.SecretKey,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "session.expire",
			Aliases:     []string{"se"},
			EnvVars:     []string{"SESSION_EXPIRE"},
			Usage:       "Session expire time in seconds",
			Value:       604800, // 7 days
			Destination: &config.Session.Expire,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.issuer",
			Aliases:     []string{"si"},
			EnvVars:     []string{"SESSION_ISSUER"},
			Usage:       "Session issuer",
			Value:       "zotion",
			Destination: &config.Session.Issuer,
		}),
	}
}
