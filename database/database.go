package database

import (
	"context"
	"fmt"
	"time"

	"github.com/FiberApps/core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func Connect(url string, database string) {
	log := logger.NewLogger()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info(fmt.Printf("Connected to Database::[%s]", database))
	DB = *client.Database(database)
	defer cancel()
}
