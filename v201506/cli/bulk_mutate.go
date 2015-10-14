package main

import (
	//"encoding/json"
	"flag"
	"crypto/rand"
	"fmt"
	gads "github.com/bbachtel/gads/v201506"
	"log"
	"time"
)

var configJson = flag.String("oauth", "./oauth.json", "API credentials")

// Your campaign ID should go here
var campaignId int64 = 211793582

func main() {
	flag.Parse()
	config, err := gads.NewCredentialsFromFile(*configJson)
	if err != nil {
		log.Fatal(err)
	}
	// Bulk Mutate
	ms := gads.NewMutateJobService(&config.Auth)
	policy := new(gads.BulkMutateJobPolicy)

	// If you need to add prerequisites
	//policy.PrerequisiteJobIds = append(policy.PrerequisiteJobIds, 123456)

	ago := gads.AdGroupOperations{
				"ADD": {
					gads.AdGroup{
						Name:       "test ad group " + rand_str(10),
						Status:     "PAUSED",
						CampaignId: campaignId,
					},
					gads.AdGroup{
						Name:       "test ad group " + rand_str(10),
						Status:     "PAUSED",
						CampaignId: campaignId,
					},
				},
			}

	if resp, err := ms.Mutate(ago, policy); err == nil {
		jobId := resp.Id

		// loop
		for {
			// recheck every 5 seconds
			time.Sleep(5 * time.Second)

			jobSelector := gads.BulkMutateJobSelector{JobIds: []int64{jobId},Xsi_type: "BulkMutateJobSelector"}

			result, err := ms.Get(jobSelector)

			if err != nil {
				panic(err)
			}

			fmt.Println(result)
			break
		}
	} else {
		// handle err
		panic(err)
	}


}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}