package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Mongo struct {
	ConvStartMsg    string `bson:"conv_start_msg"`
	BusinessDetails []struct {
		BusinessName string `bson:"business_name"`
		Location     string `bson:"location"`
		PhoneNumber  string `bson:"phone_number"`
	} `bson:"business_details"`
}

type Business struct {
	BusinessName string `bson:"business_name"`
	Location     string `bson:"location"`
	PhoneNumber  string `bson:"phone_number"`
}

type UpdatedBusiness struct {
	BusinessName string `bson:"business_name"`
	Location     string `bson:"location"`
	PhoneNumber  string `bson:"phone_number"`
	ID           int
	NumLeads     int
	NumResponse  int
	NumAccepts   int
}

func main() {
	fmt.Println("userSearchFiltering: accessing environment variable:- mongoUrl")
	mongoUrl := os.Getenv("mongoUrl")
	// mongoUrl = "mongodb://developer:thedatabaseboss@cluster0.txuwl.mongodb.net/Bino_search?retryWrites=true&w=majority"
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

	database := client.Database("Bino_search")
	fmt.Printf("Database name: %s\n", database.Name())
	condition := bson.M{
		"state":            "business_send",
		"business_details": bson.M{"$exists": true, "$not": bson.M{"$size": 0}},
	}

	// options to the search query of mongo
	// findOptions := options.Find()
	// findOptions.SetLimit(1000)

	cur, err := database.Collection("users").Find(context.Background(), condition)
	if err != nil {
		log.Fatal(err)
	}

	var data []Mongo
	if err := cur.All(context.Background(), &data); err != nil {
		log.Fatal(err)
	}

	// disconnect from mongo
	func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}()

	// var output []bson.M
	output := make(map[string]Business)
	// Use Regex to match for travel, travelling etc
	regex := `(?i)\btravel(?:ling|led)?\b`

	for _, val := range data {
		if ok, _ := regexp.Match(regex, []byte(val.ConvStartMsg)); ok {
			businessList := val.BusinessDetails
			for _, business := range businessList {
				phone := business.PhoneNumber
				_, ok := output[phone]
				if !ok {
					output[phone] = business
				}
			}
		}
	}

	fmt.Printf("Num output %v\n", len(output))
	total := len(output)
	// Call bowApi
	updatedBusinessMap := make(map[string]UpdatedBusiness)
	fmt.Println("Calling bowApi to get business_id from phoneNum")
	flag := 0

	// add progress bar
	bar := progressbar.Default(int64(total))
	for number := range output {
		id := getBowIDFromPhoneNum(number)
		if id == 0 {
			continue
		}
		// fmt.Println(id)
		leads := getCountFromMongo("Bino_search", "broadcasts", "business_user_id", strconv.Itoa(id))
		leadsRepliedTo := getCountFromMongo("Bino_search", "convreplies", "replied_business_id", strconv.Itoa(id))
		leadsAccepted := getLeadsAccepted(id)

		var temp UpdatedBusiness
		temp.BusinessName = output[number].BusinessName
		temp.Location = output[number].Location
		temp.PhoneNumber = output[number].PhoneNumber
		temp.ID = id
		temp.NumLeads = leads
		temp.NumResponse = leadsRepliedTo
		temp.NumAccepts = leadsAccepted

		updatedBusinessMap[number] = temp
		// fmt.Printf("leads: %v\n", leads)
		// fmt.Printf("replies: %v\n", leadsRepliedTo)
		// fmt.Printf("accepts: %v\n", leadsRepliedTo)
		flag++
		bar.Add(1)
	}

	// Marshalling json
	out, err := json.Marshal(updatedBusinessMap)
	if err != nil {
		log.Fatalf("error marshalling")
	}

	// Write marshalled json to file
	err = os.WriteFile("output.json", out, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}

func getBowIDFromPhoneNum(phone string) (bowId int) {
	url := fmt.Sprintf("https://app.bow.chat/api/v1/accounts/2/contacts/search?q=%v", phone)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error generating request object")
		return
	}
	apiKey := os.Getenv("APIKEY")
	request.Header.Set("api_access_token", apiKey)

	client := &http.Client{}
	resp, error := client.Do(request)
	if error != nil {
		fmt.Println("error making request")
		return
	}

	// Some phone numbers are not present in the database, hence, the api would return a 400. This has to ignored.
	if resp.StatusCode != 200 {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var respStruct RespStruct
	json.Unmarshal(body, &respStruct)

	if len(respStruct.Payload) == 0 {
		return
	}
	bowId = respStruct.Payload[0].ID

	return
}

func getCountFromMongo(db, collection, key, value string) (count int) {
	mongoUrl := os.Getenv("mongoUrl")

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

	database := client.Database(db)
	condition := bson.M{
		key: value,
	}

	// options to the search query of mongo
	// findOptions := options.Find()
	// findOptions.SetLimit(1000)
	cnt, err := database.Collection(collection).CountDocuments(context.Background(), condition)
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()
	count = int(cnt)

	return
}

func getLeadsAccepted(id int) (count int) {
	mongoUrl := os.Getenv("mongoUrl")
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

	database := client.Database("Bino_search")
	condition := bson.M{
		"replied_business_id": strconv.Itoa(id),
	}

	cnt, err := database.Collection("convreplies").Find(context.Background(), condition)
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	var data []ConvReplies
	if err := cnt.All(context.Background(), &data); err != nil {
		fmt.Println("error marshalling data")
		fmt.Println(err)
		return
	}
	num := 0
	for _, reply := range data {
		for _, svReply := range reply.ConvReply {
			if svReply.ReplyContent == "Accept" {
				num++
				break
			}
		}
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()
	count = num
	return
}
