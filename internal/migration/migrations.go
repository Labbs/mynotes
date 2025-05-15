package migration

import (
	"embed"

	"github.com/labbs/mynotion/internal/logger/zerolog"

	_ "github.com/labbs/mynotion/internal/migration/files"
	"github.com/pressly/goose/v3"
	z "github.com/rs/zerolog"
	"gorm.io/gorm"
)

//go:embed files/*
var migrationFiles embed.FS

func RunMigration(l z.Logger, db *gorm.DB) error {
	goose.SetBaseFS(migrationFiles)
	logger := l.With().Str("event", "migration").Logger()
	goose.SetLogger(&zerolog.ZerologGooseAdapter{Logger: logger})

	// Set the dialect following the gorm dialect
	dbDialect := db.Dialector.Name()

	if err := goose.SetDialect(dbDialect); err != nil {
		logger.Error().Err(err).Msg("Failed to set dialect")
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get sql db")
		return err
	}

	if err := goose.Up(sqlDB, "files"); err != nil {
		if err.Error() != "no change" {
			logger.Error().Err(err).Msg("Failed to run migrations")
			return err
		}
	}

	return nil
}
