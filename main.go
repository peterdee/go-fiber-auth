package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"

	"go-fiber-auth/apis/index"
	"go-fiber-auth/configuration"
	"go-fiber-auth/database"
	"go-fiber-auth/utilities"
)

func main() {
	// load environment variables via the .env file
	env := os.Getenv("ENV")
	if env != "heroku" {
		envError := godotenv.Load()
		if envError != nil {
			log.Fatal(envError)
			return
		}
	}

	// connect to the database
	dbError := database.Connect()
	if dbError != nil {
		log.Fatal(dbError)
		return
	}

	app := fiber.New()

	// middlewares
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max:      60,
		Duration: 60 * time.Second,
		LimitReached: func(ctx *fiber.Ctx) error {
			return utilities.Response(utilities.ResponseParams{
				Ctx:    ctx,
				Info:   configuration.ResponseMessages.TooManyRequests,
				Status: fiber.StatusTooManyRequests,
			})
		},
	}))
	app.Use(logger.New())

	// available APIs
	app.Get("/", index.GetIndex)
	app.Get("/api", index.GetIndex)

	// handle 404
	app.Use(func(ctx *fiber.Ctx) error {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
	})

	// get the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9119"
	}

	// launch the app
	launchError := app.Listen(":" + port)
	if launchError != nil {
		panic(launchError)
	}
}
