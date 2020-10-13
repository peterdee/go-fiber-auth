package middlewares

import (
	"go-fiber-auth/configuration"
	"go-fiber-auth/utilities"

	"github.com/gofiber/fiber/v2"
)

// Authorize requests
func Authorize(ctx *fiber.Ctx) error {
	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.InternalServerError,
		Status: fiber.StatusInternalServerError,
	})
}
