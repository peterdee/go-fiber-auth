package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Instance MongoInstance

// Create database connection
func Connect() error {
	DatabaseConnection := os.Getenv("DATABASE_CONNECTION")
	DatabaseName := os.Getenv("DATABASE_NAME")
	client, clientError := mongo.NewClient(options.Client().ApplyURI(DatabaseConnection))
	if clientError != nil {
		return clientError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionError := client.Connect(ctx)
	db := client.Database(DatabaseName)

	if connectionError != nil {
		return connectionError
	}

	Instance = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
