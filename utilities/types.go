package utilities

import (
	JWT "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type GenerateJWTParams struct {
	ExpiresIn int64
	UserId    string
}

type JWTClaims struct {
	UserId string `json:"userId"`
	JWT.StandardClaims
}

type ResponseParams struct {
	Ctx    *fiber.Ctx
	Data   interface{}
	Info   string
	Status int
}
