package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/mynotion/pkg/models"
	"github.com/rs/zerolog"
)

type AuthController struct {
	AuthService    models.AuthService
	SessionService models.SessionService
	Logger         zerolog.Logger
}

// Login godoc
// @Summary Login
// @Description Login with email or username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login request"
// @Success 200 {object} models.LoginResponse
// @Route POST /api/auth/login
func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.auth.login").Logger()

	var loginRequest models.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		logger.Error().Err(err).Msg("Error parsing login request")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	loginResponse, err := ac.AuthService.Login(loginRequest)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting user by email or username")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Get session information for register the session
	session := new(models.Session)
	session.UserId = loginResponse.UserId
	session.Id = loginResponse.SessionId
	session.UserAgent = ctx.Get("User-Agent")
	session.IpAddress = ctx.IP()

	// Create the session in the database
	if err := ac.SessionService.Create(session); err != nil {
		logger.Error().Err(err).Msg("Error creating session")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Info().Str("user", loginRequest.Email).Str("session_id", loginResponse.SessionId).Msg("User logged in successfully")

	return ctx.Status(fiber.StatusOK).JSON(loginResponse)
}

// Logout godoc
// @Summary Logout
// @Description Logout the user
// @Tags auth
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /api/auth/logout [get]
func (ac *AuthController) Logout(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.auth.logout").Logger()

	// Get the user ID from the JWT token
	userId := ctx.Locals("user_id").(string)

	// Get the session ID from the JWT token
	sessionId := ctx.Locals("session_id").(string)
	// Delete the session from the database
	if err := ac.SessionService.Delete(sessionId); err != nil {
		logger.Error().Err(err).Msg("Error deleting session")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logger.Info().Str("user_id", userId).Msg("User logged out successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}

// Register godoc
// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body models.RegisterRequest true "Register request"
// @Success 201 {object} models.RegisterResponse
// @Failure 400 {object} fiber.Map
// @Failure 409 {object} fiber.Map
// @Router /api/auth/register [post]
func (ac *AuthController) Register(ctx *fiber.Ctx) error {
	logger := ac.Logger.With().Str("event", "api.auth.register").Logger()

	var registerRequest models.RegisterRequest
	if err := ctx.BodyParser(&registerRequest); err != nil {
		logger.Error().Err(err).Msg("Error parsing register request")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	registerResponse, err := ac.AuthService.Register(registerRequest)
	if err != nil {
		logger.Error().Err(err).Msg("Error registering user")
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}

	logger.Info().Str("user", registerRequest.Email).Msg("User registered successfully")
	return ctx.Status(fiber.StatusCreated).JSON(registerResponse)
}
