package main

import (
    "fmt"
    "os"
    "log"
	"time"
    "context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    fmt.Println("Hello World")
    mongoUrl := os.Getenv("mongoUrl")
    fmt.Println(mongoUrl)
    clientOptions := options.Client().ApplyURI(mongoUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	database := client.Database("TestDB")
	fmt.Printf("Database name: %s\n", database.Name())
    condition := bson.M{
        "status": "business_send",
    }
    fmt.Println(condition)


	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}()
}
