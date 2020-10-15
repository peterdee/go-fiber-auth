package auth

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-auth/configuration"
	. "go-fiber-auth/database"
	. "go-fiber-auth/database/schemas"
	"go-fiber-auth/utilities"
)

// Handle signing in
func signIn(ctx *fiber.Ctx) error {
	// check data
	var body SignInUserRequest
	bodyParsingError := ctx.BodyParser(&body)
	if bodyParsingError != nil {
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
	UserCollection := Instance.Database.Collection("User")

	// find a user
	rawUserRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "email", Value: trimmedEmail}},
	)
	userRecord := &User{}
	rawUserRecord.Decode(userRecord)
	if userRecord.ID == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	// load Password schema
	PasswordCollection := Instance.Database.Collection("Password")

	// find a password
	rawPasswordRecord := PasswordCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "userId", Value: userRecord.ID}},
	)
	passwordRecord := &Password{}
	rawPasswordRecord.Decode(passwordRecord)
	if passwordRecord.ID == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	// compare hashes
	passwordIsValid := utilities.CompareHashes(trimmedPassword, passwordRecord.Hash)
	if !passwordIsValid {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	accessExpiration, expirationError := strconv.Atoi(os.Getenv("TOKENS_ACCESS_EXPIRATION"))
	if expirationError != nil {
		accessExpiration = 24
	}
	token, tokenError := utilities.GenerateJWT(utilities.GenerateJWTParams{
		ExpiresIn: int64(accessExpiration),
		UserId:    userRecord.ID,
	})
	if tokenError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"token": token,
			"user":  userRecord,
		},
	})
}
