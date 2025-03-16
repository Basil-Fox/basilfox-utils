package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client

// Connect establishes a connection to the MongoDB database and verifies connectivity.
func Connect(url string, database string) error {
	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	// Ping MongoDB to verify connectivity
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	// Assign the database to the global variable
	DB = client.Database(database)

	return nil
}

// Disconnect gracefully disconnects from MongoDB.
func Disconnect() error {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return client.Disconnect(ctx)
	}
	return nil
}
