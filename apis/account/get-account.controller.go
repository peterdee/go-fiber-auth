package account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-auth/configuration"
	. "go-fiber-auth/database"
	. "go-fiber-auth/database/schemas"
	"go-fiber-auth/utilities"
)

// Get own account
func getAccount(ctx *fiber.Ctx) error {
	// get User ID
	stringId := fmt.Sprintf("%v", ctx.Locals("UserId"))
	userId, primitiveError := primitive.ObjectIDFromHex(stringId)
	if primitiveError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// find user record
	rawUserRecord := Instance.Database.Collection("User").FindOne(
		ctx.Context(),
		bson.D{{Key: "_id", Value: userId}},
	)
	userRecord := &User{}
	decodeError := rawUserRecord.Decode(userRecord)
	if userRecord.ID == "" || decodeError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.AccessDenied,
			Status: fiber.StatusUnauthorized,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"user": userRecord,
		},
	})
}
