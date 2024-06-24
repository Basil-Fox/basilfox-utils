package database

import (
	"context"
	"time"

	"github.com/FiberApps/common-library/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func Connect(url string, database string) {
	log := logger.New()

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Error("DATABASE:: Error creating new db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Error("DATABASE:: Error obtaining context: %v", err)
	}

	log.Info("DATABASE:: Connected to database: %s", database)
	DB = *client.Database(database)
	defer cancel()
}
