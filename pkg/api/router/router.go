package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/middleware/rbac"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
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

	crbac := rbac.Config{
		Logger:          c.Logger,
		UserService:     service.NewUserService(ur),
		GroupService:    service.NewGroupService(gr),
		SpaceService:    service.NewSpaceService(ssr),
		DocumentService: service.NewDocumentService(dr),
	}

	NewAuthRouter(c, crbac.Check())
	NewMeRouter(c, crbac.Check())
	NewDocumentRouter(c, crbac.Check())
	NewAdminRouter(c, crbac.Check())
	NewPublicRouter(c, crbac.Check())
	NewSpaceRouter(c, crbac.Check())
}
