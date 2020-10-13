package utilities

import "github.com/gofiber/fiber/v2"

type ResponseParams struct {
	Ctx    *fiber.Ctx
	Data   interface{}
	Info   string
	Status int
}

type GenerateJWTParams struct {
	ExpiresIn int64
	UserId    string
}
