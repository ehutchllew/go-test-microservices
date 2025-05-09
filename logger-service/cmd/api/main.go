package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	GRPC_PORT = 50002
	MONGO_URL = "mongodb://mongo:27017"
	PORT      = 9002
	RPC_PORT  = 5002
)

var client *mongo.Client

type Config struct{}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient
}

func connectToMongo() (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(MONGO_URL)
	clientOpts.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Println("Error connectin:", err)
		return nil, err
	}

	return c, nil
}
