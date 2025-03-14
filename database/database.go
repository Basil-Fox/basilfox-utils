package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

// Connect establishes a connection to the MongoDB database and verifies connectivity.
func Connect(url string, database string) error {
	// Create a new client and connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB client
	if err := client.Connect(ctx); err != nil {
		return err
	}

	// Ping MongoDB to verify connectivity
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	// Assign the database to the DB variable
	DB = *client.Database(database)

	// Ensure the client is disconnected when done
	// Returning an error handler to ensure proper cleanup
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			// You can log this error if needed, but avoid returning here to avoid overwriting the original error
			// log.Printf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	return nil
}
