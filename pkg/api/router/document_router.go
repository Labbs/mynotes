package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/middleware"
	"github.com/labbs/zotion/pkg/api/v1/controller"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewDocumentRouter(config *Config, rbacMiddleware fiber.Handler) {
	// Set up the document routes
	config.Logger.Info().Msg("Setting up document routes")

	// initialize the document repository
	dr := repository.NewDocumentRepository(config.Db)

	// initialize the user repository with the database connection
	c := controller.DocumentController{
		DocumentService: service.NewDocumentService(dr),
		Logger:          config.Logger,
	}

	v1Document := config.Fiber.Group("/api/v1/document", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(repository.NewSessionRepository(config.Db))), rbacMiddleware)
	v1Document.Get("/space/:spaceId", c.GetDocumentsFromSpace)
	v1Document.Get("/space/:spaceId/parent/:documentId", c.GetDocumentsFromParentDocument)
	v1Document.Get("/slug/:slug", c.GetDocumentBySlug)
	v1Document.Get("/:documentId", c.GetDocumentById)
	v1Document.Post("/", c.CreateDocument)
	v1Document.Put("/:documentId", c.UpdateDocument)
	v1Document.Delete("/:documentId", c.DeleteDocument)
}
