package v201609

import "encoding/xml"

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
type AdGroupExtensionSettingService struct {
	Auth
}

func NewAdGroupExtensionSettingService(auth *Auth) *AdGroupExtensionSettingService {
	return &AdGroupExtensionSettingService{Auth: *auth}
}

// This field can be selected using the value "Extensions".
type ExtensionType int

const (
	ExtensionFeedItemType ExtensionType = iota
	AppFeedItemType
	CallFeedItemType
	CalloutFeedItemType
	PriceFeedItemType
	ReviewFeedItem
	SitelinkFeedItemType
	StructuredSnippetFeedItemType
)

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
func (s *AdGroupExtensionSettingService) Query(query string) (settings []AdGroupExtensionSetting4Call, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		adGroupExtensionSettingServiceUrl,
		"query",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "query",
			},
			Query: query,
		},
	)

	if err != nil {
		return
	}

	getResp := struct {
		Size     int64                          `xml:"rval>totalNumEntries"`
		Settings []AdGroupExtensionSetting4Call `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return
	}
	return getResp.Settings, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.AdGroupExtensionSetting
// An AdGroupExtensionSetting is used to add or modify extensions being served for the specified ad group.
type AdGroupExtensionSetting4Call struct {
	AdGroupId        int64                 `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupId,omitempty"`
	ExtensionType    FeedType              `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting ExtensionSetting4Call `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}
