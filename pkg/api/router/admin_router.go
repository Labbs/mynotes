package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/v1/controller"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewAdminRouter(config *Config, rbacMiddleware fiber.Handler) {
	// Set up the admin routes
	config.Logger.Info().Msg("Setting up admin routes")

	// initialize the user repository
	ur := repository.NewUserRepository(config.Db)
	gr := repository.NewGroupRepository(config.Db)
	sr := repository.NewSpaceRepository(config.Db)
	dr := repository.NewDocumentRepository(config.Db)

	// initialize the user repository with the database connection
	c := controller.AdminController{
		UserService:     service.NewUserService(ur),
		GroupService:    service.NewGroupService(gr),
		SpaceService:    service.NewSpaceService(sr),
		DocumentService: service.NewDocumentService(dr),
		Logger:          config.Logger,
	}

	v1Admin := config.Fiber.Group(ApiV1Path+"/admin", rbacMiddleware)
	v1Admin.Get("/users", c.GetUsers)
	v1Admin.Get("/groups", c.GetGroups)
	v1Admin.Get("/spaces", c.GetSpaces)
}
