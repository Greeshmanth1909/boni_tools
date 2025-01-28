package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var ImagePathsArray []string = []string{
	"http://photos.hotelbeds.com/giata/small/%v",    // 74px wide (thumbnail)
	"http://photos.hotelbeds.com/giata/medium/%v",   // 117px wide
	"http://photos.hotelbeds.com/giata/bigger/%v",   // 800px wide
	"http://photos.hotelbeds.com/giata/xl/%v",       // 1024px wide
	"http://photos.hotelbeds.com/giata/xxl/%v",      // 2048px wide
	"http://photos.hotelbeds.com/giata/original/%v", // Original size
	"http://photos.hotelbeds.com/giata/%v",          // Default size
}

func main() {
	fmt.Println("Starting Content Api Fetching")
	// Upload()
	// return
	apiKey := os.Getenv("API")
	apiSecret := os.Getenv("SECRET")
	mongoUrl := os.Getenv("MONGOURL")
	fmt.Println("Mongo URL: ", mongoUrl)
	fmt.Println("Starting Content Api Fetching")

	var output HotelsDB
	start := 1 // start from 10001
	end := 1000
	for end <= 10000 {
		response := MakeRquestWithRange(apiKey, apiSecret, start, end)
		start = end + 1
		end = end + 1000
		for _, hotel := range response.Hotels {
			for i := range hotel.Images {
				path := hotel.Images[i].Path
				hotel.Images[i].ImagePaths = GenerateImagePaths(path)
			}
			output.Hotels = append(output.Hotels, hotel)
		}
		time.Sleep(1 * time.Second)
	}
	// for _, hotel := range output.Hotels {
	// 	for _, image := range hotel.Images {
	// 		fmt.Println("Image Paths: ", image.ImagePaths)
	// 	}
	// }
	fmt.Println("fetched documents from ", start, "To ", end)
	data, _ := json.Marshal(output)
	os.WriteFile("output.json", data, 0644)
	fmt.Println("DONE")
}

func GenerateXSignature(key, secret string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data := key + secret + timestamp
	hash := sha256.Sum256([]byte(data))
	signature := hex.EncodeToString(hash[:])
	return signature
}

func GenerateImagePaths(path string) []string {
	imagePaths := []string{}
	for _, imagePath := range ImagePathsArray {
		imagePaths = append(imagePaths, fmt.Sprintf(imagePath, path))
	}
	return imagePaths
}

func MakeRquestWithRange(apiKey, apiSecret string, start, end int) (response ResponseStruct) {
	xSignature := GenerateXSignature(apiKey, apiSecret)

	// Create a new HTTP client
	client := &http.Client{}

	// Create the request
	req, err := http.NewRequest("GET", "https://api.test.hotelbeds.com/hotel-content-api/1.0/hotels", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Add query parameters
	q := req.URL.Query()
	q.Add("fields", "all")
	q.Add("language", "ENG")
	q.Add("from", strconv.Itoa(start))
	q.Add("to", strconv.Itoa(end))
	q.Add("useSecondaryLanguage", "false")
	req.URL.RawQuery = q.Encode()

	// Add headers
	req.Header.Add("Api-key", apiKey)
	req.Header.Add("X-Signature", xSignature)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the response into our struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return
	}

	fmt.Printf("Total hotels found: %d\n", response.Total)
	fmt.Printf("Process Time: %s\n", response.AuditData.ProcessTime)

	return response
}
