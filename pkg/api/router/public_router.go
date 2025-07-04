package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/v1/controller"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewPublicRouter(config *Config, rbacMiddleware fiber.Handler) {
	// Set up the public routes
	config.Logger.Info().Msg("Setting up public routes")

	// initialize the document repository
	dr := repository.NewDocumentRepository(config.Db)

	// initialize the document service
	ds := service.NewDocumentService(dr)

	// initialize the document controller
	dc := controller.DocumentController{
		DocumentService: ds,
		Logger:          config.Logger,
	}

	// Set up the public routes
	public := config.Fiber.Group(ApiV1Path + "/public")
	public.Get("/document/slug/:slug", dc.GetDocumentBySlug)
}
