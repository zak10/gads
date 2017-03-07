package v201609

import (
	"encoding/xml"
	"fmt"
)

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
type AdGroupExtensionSettingService struct {
	Auth
}

func NewAdGroupExtensionSettingService(auth *Auth) *AdGroupExtensionSettingService {
	return &AdGroupExtensionSettingService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.AdGroupExtensionSetting
// An AdGroupExtensionSetting is used to add or modify extensions being served for the specified ad group.
type AdGroupExtensionSetting struct {
	AdGroupId        int64             `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupId,omitempty"`
	ExtensionType    FeedType          `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting AdGroupExtSetting `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}

type AdGroupExtensionSettingOperations map[string][]AdGroupExtensionSetting

type AdGroupExtSetting ExtensionSetting

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
func (s *AdGroupExtensionSettingService) Query(query string) (settings []AdGroupExtensionSetting, totalCount int64, err error) {
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
		Size     int64                     `xml:"rval>totalNumEntries"`
		Settings []AdGroupExtensionSetting `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return
	}
	return getResp.Settings, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#mutate
func (s *AdGroupExtensionSettingService) Mutate(settingsOperations AdGroupExtensionSettingOperations) (settings []AdGroupExtensionSetting, err error) {
	type settingOperations struct {
		Action  string                  `xml:"operator"`
		Setting AdGroupExtensionSetting `xml:"operand"`
	}
	operations := []settingOperations{}
	for action, settings := range settingsOperations {
		for _, setting := range settings {
			operations = append(operations,
				settingOperations{
					Action:  action,
					Setting: setting,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []settingOperations `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}

	respBody, err := s.Auth.request(adGroupExtensionSettingServiceUrl, "mutate", mutation)
	if err != nil {
		return settings, err
	}
	mutateResp := struct {
		Settings []AdGroupExtensionSetting `xml:"rval>value"`
	}{}
	err = xml.Unmarshal(respBody, &mutateResp)
	if err != nil {
		return settings, err
	}

	return mutateResp.Settings, err
}

func (s AdGroupExtSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if s.PlatformRestrictions != "NONE" {
		e.EncodeElement(&s.PlatformRestrictions, xml.StartElement{Name: xml.Name{
			"https://adwords.google.com/api/adwords/cm/v201609",
			"platformRestrictions"}})
	}
	switch extType := s.Extensions.(type) {
	case []CallFeedItem:
		e.EncodeElement(s.Extensions.([]CallFeedItem), xml.StartElement{
			xml.Name{baseUrl, "extensions"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "CallFeedItem"},
			},
		})
	default:
		return fmt.Errorf("unknown extension type %#v\n", extType)

	}

	e.EncodeToken(start.End())
	return nil
}

func (s *AdGroupExtSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	s.Extensions = []interface{}{}

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
