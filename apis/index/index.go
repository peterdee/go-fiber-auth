package index

import (
	"github.com/gofiber/fiber/v2"
)

// APIs setup
func Setup(app *fiber.App) {
	group := app.Group("/")

	group.Get("/", getIndex)
	group.Get("/api", getIndex)
}
