package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-auth/configuration"
	. "go-fiber-auth/database"
	. "go-fiber-auth/database/schemas"
	"go-fiber-auth/utilities"
)

// Handle signing up
func signUp(ctx *fiber.Ctx) error {
	// check data
	var body CreateUserRequest
	bodyParsingError := ctx.BodyParser(&body)
	if bodyParsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	email := body.Email
	name := body.Name
	password := body.Password
	role := body.Role
	if email == "" || name == "" || password == "" || role == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}
	trimmedEmail := strings.TrimSpace(email)
	trimmedName := strings.TrimSpace(name)
	trimmedPassword := strings.TrimSpace(password)
	trimmedRole := strings.TrimSpace(role)
	if trimmedEmail == "" || trimmedName == "" ||
		trimmedPassword == "" || trimmedRole == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}

	// make sure that the role is correct
	roles := utilities.Values(configuration.Roles)
	if !utilities.IncludesString(roles, trimmedRole) {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}

	// load User schema
	UserCollection := Instance.Database.Collection("User")

	// check if email is already in use
	existingRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "email", Value: trimmedEmail}},
	)
	existingUser := &User{}
	existingRecord.Decode(existingUser)
	if existingUser.ID != "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.EmailAlreadyInUse,
			Status: fiber.StatusBadRequest,
		})
	}

	// create a new User record, insert it and get back the ID
	now := utilities.MakeTimestamp()
	NewUser := new(User)
	NewUser.Created = now
	NewUser.Email = trimmedEmail
	NewUser.ID = ""
	NewUser.Name = trimmedName
	NewUser.Role = trimmedRole
	NewUser.Updated = now
	insertionResult, insertionError := UserCollection.InsertOne(ctx.Context(), NewUser)
	if insertionError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	createdRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "_id", Value: insertionResult.InsertedID}},
	)
	createdUser := &User{}
	createdRecord.Decode(createdUser)

	// load Password schema
	PasswordCollection := Instance.Database.Collection("Password")

	// create password hash
	hash, hashError := utilities.MakeHash(trimmedPassword)
	if hashError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// create a new Password record and insert it
	NewPassword := new(Password)
	NewPassword.Created = now
	NewPassword.Hash = hash
	NewPassword.ID = ""
	NewPassword.Updated = now
	NewPassword.UserId = createdUser.ID
	_, insertionError = PasswordCollection.InsertOne(ctx.Context(), NewPassword)
	if insertionError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"user": createdUser,
		},
	})
}
