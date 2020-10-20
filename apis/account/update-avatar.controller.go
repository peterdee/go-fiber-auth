package account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"go-fiber-auth/configuration"
	// . "go-fiber-auth/database"
	// . "go-fiber-auth/database/schemas"
	"go-fiber-auth/utilities"
)

// Update account avatar
func updateAvatar(ctx *fiber.Ctx) error {
	// get User ID and convert it to the ObjectID format
	// stringId := fmt.Sprintf("%v", ctx.Locals("UserId"))
	// userId, primitiveError := primitive.ObjectIDFromHex(stringId)
	// if primitiveError != nil {
	// 	return utilities.Response(utilities.ResponseParams{
	// 		Ctx:    ctx,
	// 		Info:   configuration.ResponseMessages.InternalServerError,
	// 		Status: fiber.StatusInternalServerError,
	// 	})
	// }

	// get avatar image
	file, fileError := ctx.FormFile("avatar")
	if fileError != nil {
		fmt.Println(fileError)
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}
	fmt.Println("this is a filename", fmt.Sprintf("./%s", file.Filename))

	// Save file to root directory:
	ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	// // find user record
	// rawUserRecord := Instance.Database.Collection("User").FindOne(
	// 	ctx.Context(),
	// 	bson.D{{Key: "_id", Value: userId}},
	// )
	// userRecord := &User{}
	// decodeError := rawUserRecord.Decode(userRecord)
	// if userRecord.ID == "" || decodeError != nil {
	// 	return utilities.Response(utilities.ResponseParams{
	// 		Ctx:    ctx,
	// 		Info:   configuration.ResponseMessages.AccessDenied,
	// 		Status: fiber.StatusUnauthorized,
	// 	})
	// }

	return utilities.Response(utilities.ResponseParams{
		Ctx: ctx,
	})
}
