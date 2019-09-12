package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "test"

func Con() (*mongo.Database, error) {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	//initiate connection
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}

	//Check for connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbname)

	return database, nil
}
