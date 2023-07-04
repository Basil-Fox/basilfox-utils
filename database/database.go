package database

import (
	"context"
	"time"

	"github.com/FiberApps/core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func Connect(url string, database string) {
	log := logger.New()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("DATABASE_CONNECTION_ERR::", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal("DATABASE_CONNECTION_ERR::", err)
	}

	log.Info("CONNECTED_TO_DATABASE::", database)
	DB = *client.Database(database)
	defer cancel()
}
