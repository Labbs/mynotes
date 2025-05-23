package router

import (
	"github.com/labbs/mynotes/pkg/api/middleware"
	"github.com/labbs/mynotes/pkg/api/v1/controller"
	"github.com/labbs/mynotes/pkg/repository"
	"github.com/labbs/mynotes/pkg/service"
)

func NewMeRouter(config *Config) {
	// Set up the me routes
	config.Logger.Info().Msg("Setting up me routes")

	// initialize the user repository
	ur := repository.NewUserRepository(config.Db)

	// initialize the space repository
	sr := repository.NewSpaceRepository(config.Db)

	// initialize the user service with the database connection
	us := service.NewUserService(ur)
	ss := service.NewSpaceService(sr)

	c := controller.MeController{
		UserService:  us,
		SpaceService: ss,
		Logger:       config.Logger,
	}

	v1Me := config.Fiber.Group("/api/v1/me", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(repository.NewSessionRepository(config.Db))), middleware.RBACCheckMiddleware(config.Logger))
	v1Me.Get("/profile", c.GetMyProfile)
	v1Me.Get("/favorites", c.GetMyFavorites)
	v1Me.Get("/spaces", c.GetMySpaces)
}
