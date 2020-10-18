package account

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/middlewares"
)

// APIs setup
func Setup(app *fiber.App) {
	group := app.Group("/api/account")

	group.Get("/", middlewares.Authorize, getAccount)
	group.Post("/avatar", middlewares.Authorize, updateAvatar)
}
