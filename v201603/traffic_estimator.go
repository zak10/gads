package v201603

import "encoding/xml"

type TrafficEstimatorService struct {
	Auth
}

func NewTrafficEstimatorService(auth *Auth) *TrafficEstimatorService {
	return &TrafficEstimatorService{Auth: *auth}
}

type Keyword struct {
	Id        int64  `xml:"https://adwords.google.com/api/adwords/cm/v201603 id,omitempty"`
	Text      string `xml:"https://adwords.google.com/api/adwords/cm/v201603 text,omitempty"`      // Text: up to 80 characters and ten words
	MatchType string `xml:"https://adwords.google.com/api/adwords/cm/v201603 matchType,omitempty"` // MatchType:  "EXACT", "PHRASE", "BROAD"
}

type KeywordEstimateRequest struct {
	Keyword Keyword `xml:"keyword"`
}

type AdGroupEstimateRequest struct {
	KeywordEstimateRequests []KeywordEstimateRequest `xml:"keywordEstimateRequests"`
	MaxCpc                  int64                    `xml:"https://adwords.google.com/api/adwords/cm/v201603 maxCpc>microAmount"`
}

type CampaignEstimateRequest struct {
	AdGroupEstimateRequests []AdGroupEstimateRequest `xml:"adGroupEstimateRequests"`
}

type TrafficEstimatorSelector struct {
	CampaignEstimateRequests []CampaignEstimateRequest `xml:"campaignEstimateRequests"`
}

type StatsEstimate struct {
	AverageCpc        int64   `xml:"averageCpc>microAmount"`
	AveragePosition   float64 `xml:"averagePosition"`
	ClickThroughRate  float64 `xml:"clickThroughRate"`
	ClicksPerDay      float64 `xml:"clicksPerDay"`
	ImpressionsPerDay float64 `xml:"impressionsPerDay"`
	TotalCost         int64   `xml:"totalCost>microAmount"`
}

type KeywordEstimate struct {
	Min StatsEstimate `xml:"min"`
	Max StatsEstimate `xml:"max"`
}

type AdGroupEstimate struct {
	KeywordEstimates []KeywordEstimate `xml:"keywordEstimates"`
}

type CampaignEstimate struct {
	AdGroupEstimates []AdGroupEstimate `xml:"adGroupEstimates"`
}

// Get returns an array of CampaignEstimates, holding AdGroupEstimates which
// hold KeywordEstimates, which hold the minimum and maximum values based on
// the requested keywords.
//
// Example
//
//	keywordEstimateRequests := []KeywordEstimateRequest{
//		KeywordEstimateRequest{
//			Keyword{
//				Text:      "mars cruise",
//				MatchType: "BROAD",
//			},
//		},
//		KeywordEstimateRequest{
//			Keyword{
//				Text:      "cheap cruise",
//				MatchType: "EXACT",
//			},
//		},
//		KeywordEstimateRequest{
//			Keyword{
//				Text:      "cruise",
//				MatchType: "EXACT",
//			},
//		},
//	}
//
//	adGroupEstimateRequests := []AdGroupEstimateRequest{
//		AdGroupEstimateRequest{
//			KeywordEstimateRequests: keywordEstimateRequests,
//			MaxCpc:                  1000000,
//		},
//	}
//
//	campaignEstimateRequests := []CampaignEstimateRequest{
//		CampaignEstimateRequest{
//			AdGroupEstimateRequests: adGroupEstimateRequests,
//		},
//	}
//
//	estimates, err := trafficEstimatorService.Get(TrafficEstimatorSelector{
//		CampaignEstimateRequests: campaignEstimateRequests,
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	for _, estimate := range estimates {
//		for _, adGroupEstimate := range estimate.AdGroupEstimates {
//			for _, keywordEstimate := range adGroupEstimate.KeywordEstimates {
//				fmt.Printf("Avg cpc: %d", keywordEstimate.Min.AverageCpc)
//			}
//		}
//	}
//
// Relevant documentation
//
// 		https://developers.google.com/adwords/api/docs/reference/v201603/TrafficEstimatorService#get
//
func (s *TrafficEstimatorService) Get(selector TrafficEstimatorSelector) (res []CampaignEstimate, err error) {

	respBody, err := s.Auth.request(
		trafficEstimatorServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     TrafficEstimatorSelector `xml:"selector"`
		}{
			XMLName: xml.Name{
				Space: baseTrafficUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)

	if err != nil {
		return res, err
	}

	getResp := struct {
		CampaignEstimates []CampaignEstimate `xml:"rval>campaignEstimates"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return res, err
	}

	return getResp.CampaignEstimates, err
}
