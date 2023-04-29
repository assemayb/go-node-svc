package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         context.Context
	mongoClient *mongo.Client
	err         error
)

func init() {
}

func ConnectDB() *mongo.Client {
	ctx = context.TODO()

	// mongoConnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoConnection := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConnection)

	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	log.Printf(" ==== Connected to MongoDB ====")
	return mongoClient
}
