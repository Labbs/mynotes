package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/middleware"
	"github.com/labbs/zotion/pkg/api/v1/controller"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewSpaceRouter(config *Config, rbacMiddleware fiber.Handler) {
	// Set up the space routes
	config.Logger.Info().Msg("Setting up space routes")

	// initialize the space repository
	sr := repository.NewSpaceRepository(config.Db)

	// initialize the space service
	s := service.NewSpaceService(sr)

	// initialize the space controller
	sc := controller.SpaceController{
		SpaceService: s,
		Logger:       config.Logger,
	}

	// Set up the space routes
	space := config.Fiber.Group(ApiV1Path+"/space", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(repository.NewSessionRepository(config.Db))))
	space.Post("/", sc.CreateSpace)
}
