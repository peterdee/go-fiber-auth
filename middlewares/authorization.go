package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/configuration"
	"go-fiber-auth/utilities"
)

// Authorize requests
func Authorize(ctx *fiber.Ctx) error {
	// get authorization header
	rawToken := ctx.Get("Authorization")
	if rawToken == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingToken,
			Status: fiber.StatusUnauthorized,
		})
	}
	trimmedToken := strings.TrimSpace(rawToken)
	if trimmedToken == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingToken,
			Status: fiber.StatusUnauthorized,
		})
	}

	// parse JWT
	claims, parsingError := utilities.ParseClaims(trimmedToken)
	if parsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	// store User ID in Locals so that it can be accessed later and proceed
	ctx.Locals("UserId", claims.UserId)
	return ctx.Next()
}
