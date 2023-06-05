package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(uri string, user string, password string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: user,
		Password: password,
		DB:       0, // use default DB
	})
	defer Client.Close()

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to Redis Server")
}
