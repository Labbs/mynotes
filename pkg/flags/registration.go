package flags

import (
	"github.com/labbs/zotion/pkg/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// RegistrationFlags returns a slice of cli.Flag for registration configuration.
// It's used to set up registration-related flags for the CLI application.
func RegistrationFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "registration.enabled",
			Aliases:     []string{"re"},
			EnvVars:     []string{"REGISTRATION_ENABLED"},
			Usage:       "Enable user registration",
			Value:       true,
			Destination: &config.Registration.Enabled,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "registration.require-email-verification",
			Aliases:     []string{"rev"},
			EnvVars:     []string{"REGISTRATION_REQUIRE_EMAIL_VERIFICATION"},
			Usage:       "Require email verification for new registrations",
			Value:       true,
			Destination: &config.Registration.RequireEmailVerification,
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:        "registration.domain-whitelist",
			Aliases:     []string{"rdw"},
			EnvVars:     []string{"REGISTRATION_DOMAIN_WHITELIST"},
			Usage:       "List of allowed domains for registration",
			Value:       &cli.StringSlice{},
			Destination: &config.Registration.DomainWhitelist,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "registration.password-min-length",
			Aliases:     []string{"rpl"},
			EnvVars:     []string{"REGISTRATION_PASSWORD_MIN_LENGTH"},
			Usage:       "Minimum password length for registration",
			Value:       12,
			Destination: &config.Registration.PasswordMinLength,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "registration.password-complexity",
			Aliases:     []string{"rpc"},
			EnvVars:     []string{"REGISTRATION_PASSWORD_COMPLEXITY"},
			Usage:       "Require complex passwords (uppercase, lowercase, numbers, symbols)",
			Value:       true,
			Destination: &config.Registration.PasswordComplexity,
		}),
	}
}
