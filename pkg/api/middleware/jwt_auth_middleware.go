package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/internal/tokenutil"
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

func JwtAuthMiddleware(logger zerolog.Logger, sessionService models.SessionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_logger := logger.With().Str("request_id", fmt.Sprintf("%v", c.Locals("requestid"))).Logger()

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			_logger.Error().Str("event", "middleware.jwt_auth_middleware.missing_authorization_header").Msg("Missing authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing authorization header",
			})
		}
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authorized, err := tokenutil.IsAuthorized(t[1])
			if authorized {
				user_id, sessionId, err := tokenutil.GetSessionInformationFromToken(t[1])
				if err != nil {
					_logger.Error().Err(err).Str("event", "middleware.jwt_auth_middleware.get_session_id_from_token").Msg("Error getting session id from token")
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Invalid authorization token",
					})
				}

				// Check if the session is valid
				session, err := sessionService.GetById(sessionId)
				if err != nil {
					_logger.Error().Err(err).Str("event", "middleware.jwt_auth_middleware.get_session_by_id").Msg("Error getting session by id")
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Invalid authorization token",
					})
				}

				if session.UserId != user_id {
					_logger.Error().Str("event", "middleware.jwt_auth_middleware.session_user_id_mismatch").Msg("Session user id mismatch")
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Invalid authorization token",
					})
				}

				c.Context().SetUserValue("session_id", sessionId)
				c.Context().SetUserValue("user_id", user_id)
				return c.Next()
			}
			_logger.Error().Err(err).Str("event", "middleware.jwt_auth_middleware.is_authorized").Msg("Error checking if token is authorized")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid authorization token",
			})
		}
		_logger.Error().Str("event", "middleware.jwt_auth_middleware.invalid_authorization_header").Msg("Invalid authorization header")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authorization header",
		})
	}
}
