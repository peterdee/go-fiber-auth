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
	client, err := mongo.NewClient(options.Client().ApplyURI(DatabaseConnection))

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(DatabaseName)

	if err != nil {
		return err
	}

	Instance = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
