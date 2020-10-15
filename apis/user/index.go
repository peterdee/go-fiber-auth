package user

import (
	"go-fiber-auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

// APIs setup
func Setup(app *fiber.App) {
	group := app.Group("/api/user")

	group.Get("/", middlewares.Authorize, myAccount)
}
