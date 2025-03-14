package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

// Connect establishes a connection to the Redis server.
func Connect(uri, user, password string) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: user,
		Password: password,
		DB:       0, // use default DB
	})

	// Create a background context for Ping
	ctx := context.Background()

	// Ping the Redis server to check the connection
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis server: %w", err)
	}

	return nil
}

// Close gracefully shuts down the Redis client.
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}
