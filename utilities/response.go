package utilities

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/configuration"
)

// Send a response
func Response(params ResponseParams) error {
	data := params.Data
	info := params.Info
	status := params.Status
	if info == "" {
		info = configuration.ResponseMessages.Ok
	}
	if status == 0 {
		status = fiber.StatusOK
	}

	// caclulate request latency
	initial := params.Ctx.Context().Time()
	latency := MakeTimestamp() - (initial.UnixNano() / 1e6)

	// create a response map
	responseMap := fiber.Map{
		"datetime": MakeTimestamp(),
		"info":     info,
		"latency":  latency,
		"request":  params.Ctx.OriginalURL() + " [" + params.Ctx.Method() + "]",
		"status":   status,
	}

	if data != nil {
		responseMap["data"] = data
	}

	return params.Ctx.Status(params.Status).JSON(responseMap)
}
