package main

import (
	//"encoding/json"
	"flag"
	"crypto/rand"
	"fmt"
	gads "github.com/bbachtel/gads/v201509"
	"log"
	//"time"
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
	// Batch Job
	bs := gads.NewBatchJobService(&config.Auth)

	// If you need to add prerequisites
	//policy.PrerequisiteJobIds = append(policy.PrerequisiteJobIds, 123456)

	// Creating AdGroups
	ago := gads.AdGroupOperations{
				"ADD": {
					gads.AdGroup{
						Name:       "test brianss ad group " + rand_str(10),
						Status:     "PAUSED",
						CampaignId: campaignId,
					},
					gads.AdGroup{
						Name:       "test brianss ad group " + rand_str(10),
						Status:     "PAUSED",
						CampaignId: campaignId,
					},
				},
			}

			var operations []interface{}
			operations = append(operations, ago)

	// Updating AdGroups
	/*ago := gads.AdGroupOperations{
			"SET": {
				gads.AdGroup{
					Id: 1234567890,
					CampaignId: campaignId,
					BiddingStrategyConfiguration: []gads.BiddingStrategyConfiguration{
						gads.BiddingStrategyConfiguration{
							StrategyType: "MANUAL_CPC",
							Bids: []gads.Bid{
								gads.Bid{
									Type:   "CpcBid",
									Amount: 2000000,
								},
							},
						},
					},
				},
				gads.AdGroup{
					Id: 1234567890,
					CampaignId: campaignId,
					BiddingStrategyConfiguration: []gads.BiddingStrategyConfiguration{
						gads.BiddingStrategyConfiguration{
							StrategyType: "MANUAL_CPC",
							Bids: []gads.Bid{
								gads.Bid{
									Type:   "CpcBid",
									Amount: 2000000,
								},
							},
						},
					},
				},
			},
		}*/

	/*ago := gads.AdGroupCriterionOperations{
			"SET": {
				gads.BiddableAdGroupCriterion{
					AdGroupId: 1234567890,
					Criterion: gads.KeywordCriterion{
						Id: 1234567890,
					},
					BiddingStrategyConfiguration: &gads.BiddingStrategyConfiguration{
						Bids: []gads.Bid{
							gads.Bid{
								Type:   "CpcBid",
								Amount: 3000000,
							},
						},
					},
				},
				gads.BiddableAdGroupCriterion{
					AdGroupId: 1234567890,
					Criterion: gads.KeywordCriterion{
						Id: 1234567890,
					},
					BiddingStrategyConfiguration: &gads.BiddingStrategyConfiguration{
						Bids: []gads.Bid{
							gads.Bid{
								Type:   "CpcBid",
								Amount: 2000000,
							},
						},
					},
				},
			},
		}*/

	bjo := gads.BatchJobOperations{
		BatchJobOperations: []gads.BatchJobOperation{
			gads.BatchJobOperation{
				Operator: "ADD",
				Operand: gads.BatchJob{},
			},
		},
	}

	if resp, err := bs.Mutate(bjo); err == nil {
		jobId := resp[0].UploadUrl
		fmt.Println(jobId)
		
		bjh := gads.NewBatchJobHelper(&config.Auth)
		bjh.UploadBatchJobOperations(operations, *resp[0].UploadUrl)

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