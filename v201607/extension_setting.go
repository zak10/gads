package v201607

import (
	"encoding/xml"
	"fmt"
)

// https://developers.google.com/adwords/api/docs/reference/v201607/AdGroupExtensionSettingService.ExtensionSetting
// A setting specifying when and which extensions should serve at a given level (customer, campaign, or ad group).
type ExtensionSetting struct {
	PlatformRestrictions ExtensionSettingPlatform `xml:"platformRestrictions,omitempty"`

	Extensions Extension `xml:"https://adwords.google.com/api/adwords/cm/v201607 extensions,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201607/AdGroupExtensionSettingService.ExtensionSetting.Platform
// Different levels of platform restrictions
// DESKTOP, MOBILE, NONE
type ExtensionSettingPlatform string

type Extension interface{}

// https://developers.google.com/adwords/api/docs/reference/v201607/AdGroupExtensionSettingService.ExtensionFeedItem
// Contains base extension feed item data for an extension in an extension feed managed by AdWords.
type ExtensionFeedItem struct {
	XMLName xml.Name `json:"-" xml:"extensions"`

	FeedId                  int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201607 feedId,omitempty"`
	FeedItemId              int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201607 feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"https://adwords.google.com/api/adwords/cm/v201607 status,omitempty"`
	FeedType                *FeedType                  `xml:"https://adwords.google.com/api/adwords/cm/v201607 feedType,omitempty"`
	StartTime               string                     `xml:"https://adwords.google.com/api/adwords/cm/v201607 startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                     `xml:"https://adwords.google.com/api/adwords/cm/v201607 endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        *FeedItemDevicePreference  `xml:"https://adwords.google.com/api/adwords/cm/v201607 devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"https://adwords.google.com/api/adwords/cm/v201607 scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"https://adwords.google.com/api/adwords/cm/v201607 campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"https://adwords.google.com/api/adwords/cm/v201607 adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"https://adwords.google.com/api/adwords/cm/v201607 keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"https://adwords.google.com/api/adwords/cm/v201607 geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"https://adwords.google.com/api/adwords/cm/v201607 geoTargetingRestriction,omitempty"`
	PolicyData              *[]FeedItemPolicyData      `xml:"https://adwords.google.com/api/adwords/cm/v201607 policyData,omitempty"`

	ExtensionFeedItemType string `xml:"https://adwords.google.com/api/adwords/cm/v201607 ExtensionFeedItem.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201607/AdGroupExtensionSettingService.CallFeedItem
// Represents a Call extension.
type CallFeedItem struct {
	ExtensionFeedItem

	CallPhoneNumber               string             `xml:"https://adwords.google.com/api/adwords/cm/v201607 callPhoneNumber,omitempty"`
	CallCountryCode               string             `xml:"https://adwords.google.com/api/adwords/cm/v201607 callCountryCode,omitempty"`
	CallTracking                  bool               `xml:"https://adwords.google.com/api/adwords/cm/v201607 callTracking,omitempty"`
	CallConversionType            CallConversionType `xml:"https://adwords.google.com/api/adwords/cm/v201607 callConversionType,omitempty"`
	DisableCallConversionTracking bool               `xml:"https://adwords.google.com/api/adwords/cm/v201607 disableCallConversionTracking,omitempty"`
}

func getCallFeedItem(ext map[string]interface{}) (item CallFeedItem) {
	if val, ok := ext["CallPhoneNumber"].(string); ok {
		item.CallPhoneNumber = val
	}
	if val, ok := ext["CallCountryCode"].(string); ok {
		item.CallCountryCode = val
	}
	if val, ok := ext["CallTracking"].(bool); ok {
		item.CallTracking = val
	}
	if val, ok := ext["CallConversionType"].(CallConversionType); ok {
		if item.CallConversionType.ConversionTypeId > 0 {
			item.CallConversionType = val
		}
	}
	if val, ok := ext["DisableCallConversionTracking"].(bool); ok {
		item.DisableCallConversionTracking = val
	}
	return
}

func (s *ExtensionSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "platformRestrictions":
				if err := dec.DecodeElement(&s.PlatformRestrictions, &start); err != nil {
					return err
				}
			case "extensions":
				extension, err := extensionsUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				s.Extensions = append(s.Extensions.([]interface{}), extension)
			}
		}
	}
	return nil
}

func extensionsUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (ext interface{}, err error) {
	extensionsType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return
	}
	switch extensionsType {
	case "CallFeedItem":
		c := CallFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	default:
		err = fmt.Errorf("unknown Extensions type %#v", extensionsType)
	}
	return
}

func extensionsMarshalXML(exts []interface{}, e *xml.Encoder) error {
	for _, ext := range exts {
		var extensionType string
		extension := ext.(map[string]interface{})
		extType := FeedType(extension["FeedType"].(string))

		switch extType {
		case "CALL":
			extensionType = "CallFeedItem"
			ext = getCallFeedItem(extension)
		default:
			return fmt.Errorf("unknown extension type %#v\n", extType)
		}
		e.EncodeElement(&ext, xml.StartElement{
			xml.Name{baseUrl, "extensions"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, extensionType},
			},
		})
	}
	return nil
}
