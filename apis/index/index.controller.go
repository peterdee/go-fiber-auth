package index

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/configuration"
	"go-fiber-auth/utilities"
)

// Handle the index route
func GetIndex(ctx *fiber.Ctx) error {
	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
