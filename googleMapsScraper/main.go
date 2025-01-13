/*
    This program scrapes the google maps api and returns the data in a json file.
*/
package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "io"
    "bytes"
    "encoding/json"
)

func init() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(os.Stdout)
}

const URL string = "https://binosearch.boniserver.cloud/api/tasks/create-task-sync"

func main() {
    fmt.Println("Starting GoogleMaps Api Scraper")

    // Make a post request
    jsonStr := `{"scraper_name": "google_maps_scraper",
                    "data": {
                        "queries": ["scooter rentals in goa"],
                        "country": null,
                       "max_cities": null,
                        "randomize_cities": true,
                        "api_key": "",
                        "enable_reviews_extraction": true,
                        "max_reviews": 2,
                        "reviews_sort": "most_relevant",
                        "lang": "en",
                        "max_results": null,
                        "coordinates": "",
                        "zoom_level": 17
                    }}`

    req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(jsonStr)))
    if err != nil {
        log.Fatalf("Error generating post request\n")
    }

    req.Header.Set("Content-Type", "application/json")

    // Make the request with a http client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error getting response\n")
    }
    fmt.Println(resp.Status)

    // Read json into go struct
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Unable to read response body\n")
    }

    var responseBody Response
    json.Unmarshal(body, &responseBody)

    checkArr := make(map[string]string)
    var newResponse []Result

    // Print names of businesses to console
    for _, response := range responseBody {
        for _, val := range response.Result {
            if _, ok := checkArr[val.Phone]; !ok {
                checkArr[val.Phone] = ""
                newResponse = append(newResponse, val)
            } else {
               continue 
            }
        }
    }

    // Write to output.json
    out, err := json.Marshal(newResponse)
    if err != nil {
        log.Fatalf("Error marshalling json\n")
    }

    err = os.WriteFile("output.json", out, 0644)
    if err != nil {
        log.Fatalf("Error writing to file\n")
    }
    defer resp.Body.Close()
}
