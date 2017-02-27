package v201609

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.ExtensionSetting
// A setting specifying when and which extensions should serve at a given level (customer, campaign, or ad group).
type ExtensionSetting4Call struct {
	PlatformRestrictions ExtensionSettingPlatform `xml:"https://adwords.google.com/api/adwords/cm/v201609 platformRestrictions,omitempty"`
	Extensions           CallFeedItem             `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensions,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.ExtensionSetting.Platform
// Different levels of platform restrictions
// DESKTOP, MOBILE, NONE
type ExtensionSettingPlatform string

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.ExtensionFeedItem
// Contains base extension feed item data for an extension in an extension feed managed by AdWords.
type ExtensionFeedItem struct {
	FeedId                  int64                     `xml:"https://adwords.google.com/api/adwords/cm/v201609 feedId,omitempty"`
	FeedItemId              int64                     `xml:"https://adwords.google.com/api/adwords/cm/v201609 feedItemId,omitempty"`
	Status                  FeedItemStatus            `xml:"https://adwords.google.com/api/adwords/cm/v201609 status,omitempty"`
	FeedType                FeedType                  `xml:"https://adwords.google.com/api/adwords/cm/v201609 feedType,omitempty"`
	StartTime               string                    `xml:"https://adwords.google.com/api/adwords/cm/v201609 startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                    `xml:"https://adwords.google.com/api/adwords/cm/v201609 endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        FeedItemDevicePreference  `xml:"https://adwords.google.com/api/adwords/cm/v201609 devicePreference,omitempty"`
	Scheduling              FeedItemScheduling        `xml:"https://adwords.google.com/api/adwords/cm/v201609 scheduling,omitempty"`
	CampaignTargeting       FeedItemCampaignTargeting `xml:"https://adwords.google.com/api/adwords/cm/v201609 campaignTargeting,omitempty"`
	AdGroupTargeting        FeedItemAdGroupTargeting  `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupTargeting,omitempty"`
	KeywordTargeting        Keyword                   `xml:"https://adwords.google.com/api/adwords/cm/v201609 keywordTargeting,omitempty"`
	GeoTargeting            Location                  `xml:"https://adwords.google.com/api/adwords/cm/v201609 geoTargeting,omitempty"`
	GeoTargetingRestriction FeedItemGeoRestriction    `xml:"https://adwords.google.com/api/adwords/cm/v201609 geoTargetingRestriction,omitempty"`
	PolicyData              []FeedItemPolicyData      `xml:"https://adwords.google.com/api/adwords/cm/v201609 policyData,omitempty"`
	ExtensionFeedItemType   string                    `xml:"https://adwords.google.com/api/adwords/cm/v201609 ExtensionFeedItem.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.CallFeedItem
// Represents a Call extension.
type CallFeedItem struct {
	ExtensionFeedItem

	CallPhoneNumber               string             `xml:"https://adwords.google.com/api/adwords/cm/v201609 callPhoneNumber,omitempty"`
	CallCountryCode               string             `xml:"https://adwords.google.com/api/adwords/cm/v201609 callCountryCode,omitempty"`
	CallTracking                  bool               `xml:"https://adwords.google.com/api/adwords/cm/v201609 callTracking,omitempty"`
	CallConversionType            CallConversionType `xml:"https://adwords.google.com/api/adwords/cm/v201609 callConversionType,omitempty"`
	DisableCallConversionTracking bool               `xml:"https://adwords.google.com/api/adwords/cm/v201609 disableCallConversionTracking,omitempty"`
}
