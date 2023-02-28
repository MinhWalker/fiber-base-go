package middleware

import (
	"fiber-base-go/config"
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/services"
	"fiber-base-go/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type middlewareHandler struct {
	config      *config.Config
	userService services.UserService
}

func NewMiddlewareHandler(config *config.Config, userService services.UserService) *middlewareHandler {
	return &middlewareHandler{
		config:      config,
		userService: userService,
	}
}

func (m *middlewareHandler) RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the authorization header from the request.
		authHeader := c.Get("Authorization")

		// Verify that the authorization header is present and has the correct format.
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// Extract the token from the authorization header.
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token and extract the user's email address.
		payload, err := utils.ValidateToken(token, m.config.GoogleOAuth.ClientSecret)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// Look up the user in the database by email address.
		user := &model.User{}
		if _, err := m.userService.GetUserByEmail(payload.Email); err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		// Set the user ID in the context for downstream handlers to use.
		c.Locals("userID", user.ID)

		return c.Next()
	}
}
