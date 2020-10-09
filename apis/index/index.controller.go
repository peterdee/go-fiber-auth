package index

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/utilities"
)

// Handle the index route
func GetIndex(ctx *fiber.Ctx) error {
	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   "OK",
		Status: fiber.StatusOK,
	})
}
