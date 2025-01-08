package main

import (
    "fmt"
    "os"
    "log"
	"time"
    "context"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
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
        "state": "business_send",
        "business_details": bson.M{"$exists": true, "$not": bson.M{"$size": 0}},
    }
    fmt.Println(condition)

    findOptions := options.Find()
    findOptions.SetLimit(10)

    cur, err := database.Collection("test_collection").Find(context.Background(), condition, findOptions)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(cur)

    var data []bson.M
    if err := cur.All(context.Background(), &data); err != nil {
		log.Fatal(err)
	}

    fmt.Println(len(data))

    fmt.Println(data[5]["phone_number"])
    // Marshalling json
    out, err := json.Marshal(data)
    if err != nil {
        log.Fatalf("error marshalling")
    }

    err = os.WriteFile("output.json", out, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}()
}
