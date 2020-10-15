package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-auth/configuration"
	. "go-fiber-auth/database"
	. "go-fiber-auth/database/schemas"
	"go-fiber-auth/utilities"
)

// Get own account
func myAccount(ctx *fiber.Ctx) error {
	// load User schema
	UserCollection := Instance.Database.Collection("User")

	userId := "asd"

	// find user record
	rawUserRecord := UserCollection.FindOne(
		ctx.Context(),
		bson.D{{Key: "_id", Value: userId}},
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

	return utilities.Response(utilities.ResponseParams{
		Ctx: ctx,
		Data: fiber.Map{
			"user": userRecord,
		},
	})
}
