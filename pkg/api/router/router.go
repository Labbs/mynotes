package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotes/pkg/api/middleware"
	"github.com/labbs/mynotes/pkg/repository"
	"github.com/labbs/mynotes/pkg/service"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Config struct {
	Fiber  *fiber.App
	Logger zerolog.Logger
	Db     *gorm.DB
}

func (c *Config) Setup() {
	ur := repository.NewUserRepository(c.Db)
	gr := repository.NewGroupRepository(c.Db)
	ssr := repository.NewSpaceRepository(c.Db)
	dr := repository.NewDocumentRepository(c.Db)

	rbacMiddleware := middleware.RBACCheckMiddleware(
		c.Logger,
		service.NewUserService(ur),
		service.NewGroupService(gr),
		service.NewSpaceService(ssr),
		service.NewDocumentService(dr),
	)

	NewAuthRouter(c, rbacMiddleware)
	NewMeRouter(c, rbacMiddleware)
	NewDocumentRouter(c, rbacMiddleware)
}
