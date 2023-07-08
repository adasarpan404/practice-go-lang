package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://arpandas:hz2OaqParkUGsEsH@cluster0.jyhik.mongodb.net/cassino?retryWrites=true&w=majority"

const dbName = "netflix"

const collectionName = "movies"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return
	}

	fmt.Println("Connected to MongoDB!")

	// Close the connection when finished
	err = client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Failed to disconnect from MongoDB:", err)
		return
	}

	fmt.Println("Disconnected from MongoDB.")
}
