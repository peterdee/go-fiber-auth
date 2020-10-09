package index

import (
	"github.com/gofiber/fiber/v2"
)

// Handle the index route
func GetIndex(ctx *fiber.Ctx) error {
	// return utilities.Response(utilities.ResponseParams{
	// 	Ctx:    ctx,
	// 	Info:   configuration.ResponseMessages.Ok,
	// 	Status: fiber.StatusOK,
	// })
	return ctx.JSON(fiber.Map{
		"info": "OK",
	})
}
