package httpserver

import (
	"os"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/labbs/mynotes/internal/logger/zerolog"
	apiRouter "github.com/labbs/mynotes/pkg/api/router"
	z "github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Config struct {
	Port     int
	HttpLogs bool
	Fiber    *fiber.App
	Logger   z.Logger
	Stop     chan os.Signal
	Db       *gorm.DB
}

func (s *Config) Configure() {
	fconfig := fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	}

	r := fiber.New(fconfig)

	if s.HttpLogs {
		r.Use(zerolog.HTTPLogger(s.Logger))
	}

	// gofiber recover => https://docs.gofiber.io/api/middleware/recover
	r.Use(recover.New())

	r.Use(cors.New())
	r.Use(compress.New())
	r.Use(requestid.New())

	r.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	s.Fiber = r
}

func (s *Config) NewServer() error {
	s.Configure()

	apirc := apiRouter.Config{
		Fiber:  s.Fiber,
		Logger: s.Logger,
		Db:     s.Db,
	}

	apirc.Setup()

	go func() {
		for {
			if err := s.Fiber.Listen(":" + strconv.Itoa(s.Port)); err != nil {
				select {
				case <-s.Stop:
					return
				default:
					s.Logger.Error().Err(err).Msg("failed to listen on port")
				}
			}
		}
	}()

	s.Logger.Info().Msgf("HTTP server listening on port %d", s.Port)
	return nil
}

func (s *Config) Shutdown() error {
	s.Logger.Info().Msg("Shutting down server")
	return s.Fiber.Shutdown()
}
