package database

import (
	zerologadapter "github.com/labbs/mynotes/internal/logger/zerolog"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGorm(logger zerolog.Logger, dialect, dsn string) *gorm.DB {
	// Define the logger
	gormLogger := zerologadapter.NewGormLogger(logger)

	// Define the database variable
	var db *gorm.DB
	var err error

	// Check if the database is managed
	switch dialect {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: gormLogger})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
	default:
		logger.Fatal().Msg("Invalid database type")
	}
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	return db
}
