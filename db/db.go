package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MgoConn ...
func MgoConn() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://user:password@cluster0.c36kn.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// MgoCollection call a collection with passing client value
func MgoCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("TUTOR").Collection(coll)
}
