package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetUpMongoClient :
func SetUpMongoClient() *mongo.Client {
	var clientOptions *options.ClientOptions = options.Client().ApplyURI(Key.DBHost)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Connect :
func Connect() {
	client := SetUpMongoClient()
	// Check Connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is connected...")
}

// MongoClient :
var MongoClient *mongo.Client = SetUpMongoClient()
