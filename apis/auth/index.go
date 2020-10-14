package auth

import (
	"github.com/gofiber/fiber/v2"
)

// APIs setup
func Setup(app *fiber.App) {
	group := app.Group("/api/auth")

	group.Post("/signup", signUp)
	group.Post("/signin", signIn)
}
