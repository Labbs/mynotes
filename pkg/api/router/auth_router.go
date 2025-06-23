package router

import (
	"github.com/labbs/zotion/pkg/api/controller"
	"github.com/labbs/zotion/pkg/api/middleware"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewAuthRouter(config *Config) {
	// Set up the auth routes
	config.Logger.Info().Msg("Setting up auth routes")

	// initialize the user repository
	ur := repository.NewUserRepository(config.Db)

	// session service configuration
	sr := repository.NewSessionRepository(config.Db)

	// initialize the space repository
	ssr := repository.NewSpaceRepository(config.Db)

	// initialize the document repository
	dr := repository.NewDocumentRepository(config.Db)

	// initialize the user repository with the database connection
	c := controller.AuthController{
		AuthService:    service.NewAuthService(ur, ssr, dr),
		SessionService: service.NewSessionService(sr),
		Logger:         config.Logger,
	}

	// Set up the auth routes
	// create a new group for the auth routes
	auth := config.Fiber.Group("/api/auth")
	auth.Post("/login", c.Login)
	auth.Post("/register", c.Register)

	// create a new group for the auth routes
	// and apply the JWT authentication middleware
	// and the RBAC check middleware
	// to all routes in this group
	// this is used to protect the logout route
	// and require the user to be authenticated
	authPrivate := config.Fiber.Group("/api/auth", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(sr)), middleware.RBACCheckMiddleware(config.Logger))
	authPrivate.Post("/logout", c.Logout)
	authPrivate.Get("/validate", c.ValidateSession)
}
