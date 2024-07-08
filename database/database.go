package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func Connect(url string, database string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return err
	}

	// Ping the MongoDB to verify connectivity
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	DB = *client.Database(database)
	return nil
}
