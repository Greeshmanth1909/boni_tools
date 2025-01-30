package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/schollz/progressbar/v3"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("userSearchFiltering: accessing environment variable:- mongoUrl")
	mongoUrl := os.Getenv("mongoUrl")
	scraperUrl := os.Getenv("scraperUrl")
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

	targetMongoUrl := os.Getenv("targetMongoUrl")
	if targetMongoUrl == "" {
		fmt.Println("Please set environment variable: targetMongoUrl for the destination database")
		return
	}

	targetClientOptions := options.Client().ApplyURI(targetMongoUrl)
	targetCtx, targetCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer targetCancel()

	targetClient, err := mongo.Connect(targetCtx, targetClientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to target MongoDB: %v", err)
	}

	err = targetClient.Ping(targetCtx, nil)
	if err != nil {
		log.Fatalf("Failed to ping target MongoDB: %v", err)
	}
	fmt.Println("Connected to target MongoDB!")

	// Get data from mongo
	condition := bson.M{
		"state":            "business_send",
		"business_details": bson.M{"$exists": true, "$not": bson.M{"$size": 0}},
	}
	var data []Mongo
	getDataFromMongo("Bino_search", "users", condition, &data, client)
	fmt.Printf("Num data %v\n", len(data))

	output := make(map[string]Business)
	output2 := []Business{}

	// Compile regex once before the loop
	travelRegex, err := regexp.Compile(`(?i)\btravel(?:ling|led)?\b`)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	for _, val := range data {
		if travelRegex.Match([]byte(val.ConvStartMsg)) {
			businessList := val.BusinessDetails
			for _, business := range businessList {
				phone := business.PhoneNumber
				_, ok := output[phone]
				if !ok {
					output[phone] = business
					output2 = append(output2, business)
				}
			}
		}
	}

	fmt.Printf("Num output %v\n", len(output2))

	// Async pipeline begins
	initBuf := createBuffer(output2, client, targetClient, scraperUrl)
	asyncCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bowArr := createBowWorkers(asyncCtx, initBuf, 8)

	m1 := combineChannels(bowArr)
	m1Arr := createMongoWorkers(asyncCtx, m1, 8)

	m2 := combineChannels(m1Arr)
	m2Arr := createMongoWorkers2(asyncCtx, m2, 8)

	m3 := combineChannels(m2Arr)
	m3Arr := createMongoWorkers3(asyncCtx, m3, 8)

	scraperChan := combineChannels(m3Arr)
	scraperChanArr := createScraperWorkers(asyncCtx, scraperChan, 8)

	outputChannel := combineChannels(scraperChanArr)

	flag := 0
	bar := progressbar.Default(int64(len(output2)))

	for val := range outputChannel {
		flag++
		pushToMongo("async", "asyncImp", *val.result, val.targetClient)
		// fmt.Println(flag)
		// if flag == 30 {
		// 	fmt.Println("pushed 30 to mongo")
		// 	break
		// }
		bar.Add(1)
	}

	// f, err := os.OpenFile("asyncOut.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	//
	// for {
	// 	val, ok := <-outputChannel
	// 	if ok {
	// 		fmt.Println(val)
	// 	} else {
	// 		fmt.Println("channel is empty")
	// 		break
	// 	}
	// }

	// Async pipeline ends

	defer func() {
		if err := targetClient.Disconnect(targetCtx); err != nil {
			log.Fatalf("Failed to disconnect from target MongoDB: %v", err)
		}
		fmt.Println("Disconnected from target MongoDB!")
	}()

	// disconnect from source mongo
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}()
}

/*
getBowIDFromPhoneNum function calls the bowapi to fetch id form bowApi.
*/
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

/*
getCountFromMongo function queries the database based on the condition and returns the count of all results.
*/
func getCountFromMongo(db, collection string, condition bson.M, client *mongo.Client) (count int) {
	database := client.Database(db)
	cnt, err := database.Collection(collection).CountDocuments(context.Background(), condition)
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}
	count = int(cnt)
	return
}

/*
getLeadsAccepted function queries the db and returns the number of `Accept` that a business id has.
*/
func getLeadsAccepted(db, collection string, condition bson.M, client *mongo.Client) (count int) {
	database := client.Database(db)
	cnt, err := database.Collection(collection).Find(context.Background(), condition)
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
	count = num
	return
}

/*
getDataFromMongo function fetches data from mongo based on the given condition and stores it in results.
*/
func getDataFromMongo[T any](db, col string, condition bson.M, results *[]T, client *mongo.Client) {
	database := client.Database(db)
	fmt.Printf("Database name: %s\n", database.Name())
	cur, err := database.Collection(col).Find(context.Background(), condition)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.All(context.Background(), results); err != nil {
		log.Fatal(err)
	}
}

/*
pushToMongo function pushes the updatedBusinessMap to the specified MongoDB collection.
If the collection doesn't exist, it will be created.
*/
func pushToMongo(db, collection string, data UpdatedBusiness, client *mongo.Client) error {
	database := client.Database(db)
	col := database.Collection(collection)

	// Convert the map to an array of documents
	documents := data

	// Insert the new documents
	_, err := col.InsertOne(context.Background(), documents)
	if err != nil {
		return fmt.Errorf("error inserting documents: %v", err)
	}

	return nil
}

func createScraperTask(scraperUrl string, data string) (ScraperResponse, error, map[string]interface{}) {
	jsonStr := `{
		"scraper_name": "google_maps_scraper",
		"data": {
			"queries": [
				"%v"
			],
			"country": null,
			"max_cities": null,
			"randomize_cities": true,
			"api_key": "",
			"enable_reviews_extraction": true,
			"max_reviews": 2,
			"reviews_sort": "most_relevant",
			"lang": "en",
			"max_results": 1,
			"coordinates": "",
			"zoom_level": 17
		}
	}`

	var n map[string]interface{}
	jsonDataStrFormatted := fmt.Sprintf(jsonStr, data)

	// Convert to a byte buffer for the request
	jsonData := []byte(jsonDataStrFormatted)

	request, err := http.NewRequest("POST", scraperUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request")
		return ScraperResponse{}, nil, n
	}
	// add headers
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request")
		return ScraperResponse{}, nil, n
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return ScraperResponse{}, nil, n
	}

	var resp ScraperResponse
	_ = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling response")
		json.Unmarshal(body, &n)
		return ScraperResponse{}, errors.New("420"), n
	}

	return resp, nil, n
}
