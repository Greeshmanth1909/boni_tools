package main

import (
    "fmt"
	"strconv"
    "context"
    "sync"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type buffer struct {
	business     Business
	condition    bson.M
	client       *mongo.Client
	targetClient *mongo.Client
	scraperUrl   string
	result       *UpdatedBusiness
}

func createBuffer(output []Business, client, targetClient *mongo.Client, scraperUrl string) (<- chan buffer) {
    buf := make(chan buffer, 100)
    go func(){
        defer close(buf)
        for i, business := range output {
            newBuf := buffer{}
            newBuf.business = business
            newBuf.client = client
            newBuf.targetClient = targetClient
            newBuf.scraperUrl = scraperUrl
            newBuf.result = &UpdatedBusiness{}

            newBuf.result.BusinessName = business.BusinessName
            newBuf.result.Location = business.Location
            newBuf.result.PhoneNumber = business.PhoneNumber

            buf <- newBuf
            if i == 30 {
                fmt.Println("150 done!")
                break
            }
        }
    }()

    return buf
}

// function bowWorker reads from buf channel, gets bowId from phoneNumber, adds condition and pushes it to another channel
func createBowWorkers(ctx context.Context, buf <- chan buffer, num int) []<-chan buffer {
    var bufArr = [] <-chan buffer{}
    for i := 0; i < num; i++ {
        // get chan and add it to array
        func(ctx context.Context){
            newBuf := make(chan buffer, 1000)
            go func(){
                defer close(newBuf)
                // read with switch statement
                for {
                    select {
                    case <-ctx.Done():
                        return
                    case val, ok := <- buf:
                        if ok {
                            // val exists
                            id := getBowIDFromPhoneNum(val.business.PhoneNumber)

                            val.result.ID = id
                            val.condition = bson.M{
                                "business_user_id": strconv.Itoa(id),
                            }

                            // push to channel
                            newBuf <- val
                        } else {
                            return
                        }
                    }
                }
            }()
            bufArr = append(bufArr, newBuf)
        }(ctx)
    }
    return bufArr
}

func createMongoWorkers(ctx context.Context, buf <- chan buffer, num int) []<-chan buffer {
    // updates Num leads given to a particular business
    var bufArr = [] <-chan buffer{}
    for i := 0; i < num; i++ {
        // get chan and add it to array
        func(ctx context.Context){
            newBuf := make(chan buffer, 100)

            go func(){
                defer close(newBuf)
                // read with switch statement
                for {
                    select {
                    case <-ctx.Done():
                        return
                    case val, ok := <- buf:
                        if ok {
                            // use Val.Condition to get details from mongo
                            leads := getCountFromMongo("Bino_search", "broadcasts", val.condition, val.client)
                            val.result.NumLeads = leads
                            val.condition = bson.M{
                                "replied_business_id": strconv.Itoa(val.result.ID),
                            }

                            // push to channel
                            newBuf <- val
                        } else {
                            return
                        }
                    }
                }
            }()
            bufArr = append(bufArr, newBuf)
        }(ctx)
    }
    return bufArr
}

func createMongoWorkers2(ctx context.Context, buf <- chan buffer, num int) []<- chan buffer {
    // updates Num Responses given by a particular business
    var bufArr = [] <-chan buffer{}
    for i := 0; i < num; i++ {
        // get chan and add it to array
        func(ctx context.Context){
            newBuf := make(chan buffer, 100)

            go func(){
                defer close(newBuf)
                // read with switch statement
                for {
                    select {
                    case <-ctx.Done():
                        return
                    case val, ok := <- buf:
                        if ok {
                            // use Val.Condition to get details from mongo
                            numResponses := getCountFromMongo("Bino_search", "convreplies", val.condition, val.client)
                            val.result.NumResponse = numResponses
                            // next query uses the same condition, no change required.
                            // push to channel
                            newBuf <- val
                        } else {
                            return
                        }
                    }

                }
            }()
            bufArr = append(bufArr, newBuf)
        }(ctx)
    }
    return bufArr
}

func createMongoWorkers3(ctx context.Context, buf <- chan buffer, num int) []<-chan buffer {
    // updates numAccepts
    var bufArr = [] <-chan buffer{}
    for i := 0; i < num; i++ {
        // get chan and add it to array
        func(ctx context.Context){
            newBuf := make(chan buffer, 100)

            go func(){
                defer close(newBuf)
                // read with switch statement
                for {
                    select {
                    case <-ctx.Done():
                        return
                    case val, ok := <- buf:
                        if ok {
                            // use Val.Condition to get details from mongo
                            numAccepts := getLeadsAccepted("Bino_search", "convreplies", val.condition, val.client)
                            val.result.NumAccepts = numAccepts
                            // no condition required for next task
                            // push to channel
                            newBuf <- val
                        } else {
                            return
                        }
                    }
                }
            }()
            bufArr = append(bufArr, newBuf)
        }(ctx)
    }
    return bufArr
}

func createScraperWorkers(ctx context.Context, buf <- chan buffer, num int) [] <-chan buffer {
    // updates everuthing else by fetching form google maps scraper
    var bufArr = [] <-chan buffer{}
    for i := 0; i < num; i++ {
        // get chan and add it to array
        func(ctx context.Context){
            newBuf := make(chan buffer, 100)

            go func(){
                defer close(newBuf)
                // read with switch statement
                for {
                    select {
                    case <-ctx.Done():
                        return
                    case val, ok := <- buf:
                        if ok {
                            scraperResponse := createScraperTask(val.scraperUrl, val.business.Location)
                            val.result.Description = scraperResponse[0].Result[0].Description
                            val.result.Competitors = scraperResponse[0].Result[0].Competitors
                            val.result.DetailedAddress = scraperResponse[0].Result[0].DetailedAddress
                            val.result.FeaturedImage = scraperResponse[0].Result[0].FeaturedImage
                            val.result.Images = scraperResponse[0].Result[0].Images

                            // push to channel
                            newBuf <- val
                        } else {
                            return
                        }
                    }
                }
            }()
            bufArr = append(bufArr, newBuf)
        }(ctx)
    }
    return bufArr
}

// func combChans combines all given channels
func combineChannels(channels []<-chan buffer) <-chan buffer {
    // Create a combined output channel
    combined := make(chan buffer)

    // Start a goroutine for each input channel
    var wg sync.WaitGroup
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan buffer) {
            defer wg.Done()
            for buf := range c {
                combined <- buf // Forward values to the combined channel
            }
        }(ch)
    }

    // Close the combined channel once all input channels are processed
    go func() {
        wg.Wait()
        close(combined)
    }()

    return combined
}
