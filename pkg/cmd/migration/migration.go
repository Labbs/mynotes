package migration

import (
	"errors"

	"github.com/labbs/mynotion/internal/database"
	"github.com/labbs/mynotion/internal/migration"
	"github.com/labbs/mynotion/pkg/config"
	logger "github.com/labbs/mynotion/pkg/logger"
	"gorm.io/gorm"

	"github.com/labbs/mynotion/pkg/flags"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewInstance() *cli.Command {
	migrationFlags := getFlags()

	return &cli.Command{
		Name:   "migration",
		Usage:  "Run Alfred migrations",
		Flags:  migrationFlags,
		Before: altsrc.InitInputSourceWithContext(migrationFlags, altsrc.NewJSONSourceFromFlagFunc("config")),
		Action: runMigration,
	}
}

func getFlags() (list []cli.Flag) {
	list = append(list, flags.GenericFlags()...)
	list = append(list, flags.DatabaseFlags()...)
	list = append(list, flags.LoggerFlags()...)
	return
}

func runMigration(c *cli.Context) error {
	var db *gorm.DB

	l := logger.NewLogger(config.Logger.Level, config.Logger.Pretty, c.App.Version)

	if config.Database.DSN == "" {
		return errors.New("database gorm dsn is required")
	}

	db = database.NewGorm(l, config.Database.Dialect, config.Database.DSN)

	// Run migrations
	if err := migration.RunMigration(l, db); err != nil {
		return err
	}

	return nil
}
