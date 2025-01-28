package main

import (
	"fmt"
	"strconv"
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


var countFromMongoBuffer = make(chan buffer, 100)
var leadAcceptBuffer = make(chan buffer, 100)
var scraperBuffer = make(chan buffer, 100)
var pushToMongoBuffer = make(chan buffer, 100)

func bowIDWorker(bowIdBuffer <-chan buffer, wg *sync.WaitGroup) {
	for job := range bowIdBuffer {
		job.result.BusinessName = job.business.BusinessName
		job.result.Location = job.business.Location
		job.result.PhoneNumber = job.business.PhoneNumber
		id := getBowIDFromPhoneNum(job.business.PhoneNumber)
		leadsCondition := bson.M{
			"business_user_id": strconv.Itoa(id),
		}
		job.condition = leadsCondition
		job.result.ID = id
		pushToMongoBuffer <- job
	}
	wg.Done()
}

func countFromMongoWorker(countFromMongoBuffer <-chan buffer, wg *sync.WaitGroup) {
	for job := range countFromMongoBuffer {
		leads := getCountFromMongo("Bino_search", "convreplies", job.condition, job.client)
		job.result.NumLeads = leads

		job.condition = bson.M{
			"replied_business_id": strconv.Itoa(job.result.ID),
		}
		leadAcceptBuffer <- job
	}
	wg.Done()
}

func leadAcceptWorker(leadAcceptBuffer <-chan buffer, wg *sync.WaitGroup) {
	for job := range leadAcceptBuffer {
		leadsRepliedTo := getCountFromMongo("Bino_search", "convreplies", job.condition, job.client)
		leadsAccepted := getLeadsAccepted("Bino_search", "convreplies", job.condition, job.client)
		job.result.NumResponse = leadsRepliedTo
		job.result.NumAccepts = leadsAccepted
		scraperBuffer <- job
	}
	wg.Done()
}

func scraperWorker(scraperBuffer <-chan buffer, wg *sync.WaitGroup) {
	for job := range scraperBuffer {
		scraperResponse := createScraperTask(job.scraperUrl, job.result.Location)
		job.result.Description = scraperResponse[0].Result[0].Description
		job.result.Competitors = scraperResponse[0].Result[0].Competitors
		job.result.DetailedAddress = scraperResponse[0].Result[0].DetailedAddress
		job.result.FeaturedImage = scraperResponse[0].Result[0].FeaturedImage
		job.result.Images = scraperResponse[0].Result[0].Images
		pushToMongoBuffer <- job
	}
	wg.Done()
}

func pushToMongoWorker(pushToMongoBuffer <-chan buffer, wg *sync.WaitGroup) {
	for job := range pushToMongoBuffer {
		err := pushToMongo("async", "asyncImp", *job.result, job.targetClient)
		if err != nil {
			fmt.Printf("Error pushing data to target MongoDB: %v\n", err)
		} else {
			fmt.Printf("Pushed data to target MongoDB: %v\n", job.result.ID)
		}
	}
	wg.Done()
}

func populateBuffer(businesses map[string]Business, client *mongo.Client, targetClient *mongo.Client, url string) <-chan buffer {
    var bowIdBuffer = make(chan buffer, 100)
	go func() {
		defer close(bowIdBuffer)
		for _, business := range businesses {
			bowIdBuffer <- buffer{business: business, client: client, targetClient: targetClient, scraperUrl: url, result: &UpdatedBusiness{}}
		}
	}()
	return bowIdBuffer
}

func startWorkers(numWorkers int, wg *sync.WaitGroup) {
	for i := 0; i < numWorkers; i++ {
		go bowIDWorker(bowIdBuffer, wg)
		go countFromMongoWorker(countFromMongoBuffer, wg)
		go leadAcceptWorker(leadAcceptBuffer, wg)
		go scraperWorker(scraperBuffer, wg)
		go pushToMongoWorker(pushToMongoBuffer, wg)
		wg.Add(5)
	}
}

func merge(channels ...<-chan buffer) <-chan buffer {
	out := make(chan buffer)
	var wg sync.WaitGroup

	// Start a goroutine for each channel
	output := func(ch <-chan buffer) {
		defer wg.Done()
		for val := range ch {
			out <- val
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	// Close the output channel once all input channels are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
