package main

import (
	"encoding/json"
	"fmt"
	gads "github.com/colinmutter/gads/v201409"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	config, err := gads.NewCredentials(oauth2.NoContext)
	if err != nil {
		log.Fatal(err)
	}

	var pageSize int64 = 500
	var offset int64

	// show all Campaigns
	cs := gads.NewAdGroupService(&config.Auth)
	paging := gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("AdGroups\n")
	for {
		foundAdGroups, totalCount, err := cs.Query("SELECT AdGroupId, Name, Status WHERE CampaignId = '123456'")
		if err != nil {
			log.Fatal(err)
		}
		for _, adGroup := range foundAdGroups {
			adGroupJson, _ := json.MarshalIndent(adGroup, "", "  ")
			fmt.Printf("%s\n", adGroupJson)
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}

}
