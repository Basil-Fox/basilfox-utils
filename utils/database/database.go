package database

import (
	"context"
	"log"
	"time"

	"github.com/FiberApps/core/utils/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseURL  = config.Config("DATABASE_URI")
	databaseName = config.Config("DATABASE_NAME")
	DB           mongo.Database
)

func ConnectToDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to Database::[%s]", databaseName)
	DB = *client.Database(databaseName)
	defer cancel()
}
