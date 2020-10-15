package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/configuration"
	"go-fiber-auth/utilities"
)

// Authorize requests
func Authorize(ctx *fiber.Ctx) error {
	// get authorization header
	token := ctx.Get("Authorization")
	fmt.Println("auth", token)
	ctx.Locals("userId", token)

	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
