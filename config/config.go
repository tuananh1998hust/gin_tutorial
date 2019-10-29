package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbHost = os.Getenv("DBHOST")

// ConfigKey :
type ConfigKey struct {
	SecretOrKey string
	DBHost      string
}

// Key :
var Key = getAllKey()

// SecretOrKey :
var SecretOrKey = os.Getenv("SECRETORKEY")

// SetUpMongoClient :
func SetUpMongoClient() *mongo.Client {
	var clientOptions *options.ClientOptions
	if dbHost != "" {
		clientOptions = options.Client().ApplyURI(dbHost)
	} else {
		clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	}

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

func getAllKey() ConfigKey {
	var key ConfigKey
	dbHost := os.Getenv("DBHOST")
	secretOrKey := os.Getenv("SECRETORKEY")

	if dbHost == "" {
		dbHost = "mongodb://localhost:27017"
	}

	if secretOrKey == "" {
		secretOrKey = "secretOrKey"
	}

	key = ConfigKey{DBHost: dbHost, SecretOrKey: secretOrKey}

	return key
}
