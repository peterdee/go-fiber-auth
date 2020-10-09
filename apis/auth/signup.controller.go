package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-auth/configuration"
	"go-fiber-auth/utilities"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/schemas"
)

// Handle signing up
func SignUp(ctx *fiber.Ctx) error {
	// check data
	var body CreateUserRequest
	bodyParsingError := ctx.BodyParser(&body)
	if parsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	email := body.Email
	password := body.Password
	if email == "" || password == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	trimmedEmail := strings.TrimSpace(email)
	trimmedPassword := strings.TrimSpace(password)
	if trimmedEmail == "" || trimmedPassword == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}

	// load User schema
	User := Instance.Database.Collection("User")

	// check if email is already in use
	filter := bson.D{{Key: "email", Value: trimmedEmail}}
	existingRecord := User.FindOne(ctx.Context(), filter)
	existingUser := &Todo{}
	createdRecord.Decode(createdTodo)

	// load Password schema
	Password := Instance.Database.Collection("Password")

	// create a new User record
	todo := new(Todo)
	if errorParsing := ctx.BodyParser(todo); errorParsing != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
}
