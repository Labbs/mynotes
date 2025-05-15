package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Config struct {
	Fiber  *fiber.App
	Logger zerolog.Logger
	Db     *gorm.DB
}

func (c *Config) Setup() {
	NewAuthRouter(c)
	NewMeRouter(c)
	NewDocumentRouter(c)
}
