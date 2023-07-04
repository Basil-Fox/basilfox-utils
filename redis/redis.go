package redis

import (
	"context"

	"github.com/FiberApps/core/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(uri string, user string, password string) {
	log := logger.New()
	Client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: user,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Info("Connected to Redis Server")
}
