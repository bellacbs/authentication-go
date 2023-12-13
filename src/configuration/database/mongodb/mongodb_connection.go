package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_AUTH_DB = "MONGODB_AUTH_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongoDB_database := os.Getenv(MONGODB_AUTH_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		fmt.Println("error to connect")
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("error to ping")
		return nil, err
	}

	return client.Database(mongoDB_database), nil

}
