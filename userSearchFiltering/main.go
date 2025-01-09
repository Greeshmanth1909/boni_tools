package main

import (
    "fmt"
    "os"
    "log"
	"time"
    "context"
    "encoding/json"
    "regexp"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    fmt.Println("userSearchFiltering: accessing environment variable:- mongoUrl")
    mongoUrl := os.Getenv("mongoUrl")
    if mongoUrl == "" {
        fmt.Println("Please set environment variable: mongoUrl to the url of mongoDB")
        return
    }

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

    // options to the search query of mongo
    // findOptions := options.Find()
    // findOptions.SetLimit(1000)

    cur, err := database.Collection("test_collection").Find(context.Background(), condition)
    if err != nil {
        log.Fatal(err)
    }

    var data []bson.M
    if err := cur.All(context.Background(), &data); err != nil {
		log.Fatal(err)
	}


    var output []bson.M
    // Use Regex to match for travel, travelling etc
    regex := `(?i)\btravel(?:ling|led)?\b`
    for _, val := range data {
        conv_start, ok := val["conv_start_msg"]
        if ok {
            str, yay := conv_start.(string)
            if yay {
                if ok, err := regexp.Match(regex, []byte(str)); ok{
                    output = append(output, val)
                } else if err != nil {
                    fmt.Println("Failed to match regex error")
                }

            }
        }
    }

    // Marshalling json
    out, err := json.Marshal(output)
    if err != nil {
        log.Fatalf("error marshalling")
    }

    // Write marshalled json to file
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
