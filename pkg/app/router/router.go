package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/labbs/zotion/pkg/app/static"
)

type Config struct {
	Fiber  *fiber.App
	Logger zerolog.Logger
}

func (c *Config) Setup() {
	// Serve static assets and SPA routes
	static.NewStatic(c.Fiber)
}
