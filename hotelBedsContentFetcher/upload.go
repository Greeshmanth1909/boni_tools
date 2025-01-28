package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Upload() {
	// Read the JSON file
	data, err := os.ReadFile("output.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var hotelsDB HotelsDB
	if err := json.Unmarshal(data, &hotelsDB); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
	defer cancel()

	mongoURL := os.Getenv("MONGOURL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Get collection
	collection := client.Database("HotelBedsStatic").Collection("hotelbeds_static")

	// Insert documents
	for i, hotel := range hotelsDB.Hotels[4770:] {
		_, err := collection.InsertOne(ctx, hotel)
		if err != nil {
			fmt.Printf("Error inserting hotel %d: %v\n", i, err)
			continue
		}
		if i%100 == 0 {
			fmt.Printf("Inserted %d hotels\n", i)
		}
	}

	fmt.Printf("Successfully uploaded %d hotels\n", len(hotelsDB.Hotels))
}
