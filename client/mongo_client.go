package client

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents a MongoDB client with a connected database.
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// New initializes a new MongoDB client and connects to the specified database.
func New(uri, database string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Client:   client,
		Database: client.Database(database),
	}, nil
}
