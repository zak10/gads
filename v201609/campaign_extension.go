package v201609

import (
	"encoding/xml"
	"fmt"
)

type CampaignExtensionSettingService struct {
	Auth
}

func NewCampaignExtensionSettingService(auth *Auth) *CampaignExtensionSettingService {
	return &CampaignExtensionSettingService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/CampaignExtensionSettingService.CampaignExtensionSetting
// A CampaignExtensionSetting is used to add or modify extensions being served for the specified campaign.
type CampaignExtensionSetting struct {
	CampaignId       int64                    `xml:"https://adwords.google.com/api/adwords/cm/v201609 campaignId,omitempty"`
	ExtensionType    FeedType                 `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting campaignExtensionSetting `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}

type campaignExtensionSetting ExtensionSetting

func (s campaignExtensionSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"CampaignExtensionSetting",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&s.PlatformRestrictions, xml.StartElement{Name: xml.Name{
		"https://adwords.google.com/api/adwords/cm/v201609",
		"platformRestrictions"}})
	extensionsMarshalXML(s.Extensions, e)
	e.EncodeToken(start.End())
	return nil
}

func (s *campaignExtensionSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
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
				s.Extensions = append(s.Extensions, extension)
			}
		}
	}
	return nil
}

type CampaignExtensionSettingOperations map[string][]CampaignExtensionSetting

// https://developers.google.com/adwords/api/docs/reference/v201609/CampaignExtensionSettingService#query
func (s *CampaignExtensionSettingService) Query(query string) (settings []CampaignExtensionSetting, totalCount int64, err error) {
	respBody, err := s.Auth.request(
		campaignExtensionSettingUrl,
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
		Size     int64                      `xml:"rval>totalNumEntries"`
		Settings []CampaignExtensionSetting `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return
	}
	return getResp.Settings, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201609/CampaignExtensionSettingService#mutate
func (s *CampaignExtensionSettingService) Mutate(settingsOperations CampaignExtensionSettingOperations) (settings []CampaignExtensionSetting, err error) {
	type settingOperations struct {
		Action  string                   `xml:"operator"`
		Setting CampaignExtensionSetting `xml:"operand"`
	}
	operations := []settingOperations{}
	for action, settings := range settingsOperations {
		for _, setting := range settings {
			if err = identifyCampaignExtention(&setting); err != nil {
				return settings, err
			}
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

	respBody, err := s.Auth.request(campaignExtensionSettingUrl, "mutate", mutation)
	if err != nil {
		return settings, err
	}
	mutateResp := struct {
		Settings []CampaignExtensionSetting `xml:"rval>value"`
	}{}
	err = xml.Unmarshal(respBody, &mutateResp)
	if err != nil {
		return settings, err
	}

	return mutateResp.Settings, err
}

func identifyCampaignExtention(setting *CampaignExtensionSetting) (err error) {
	switch setting.ExtensionType {
	case "CALL":
		for _, ext := range setting.ExtensionSetting.Extensions {
			item := getCallFeedItem(ext.(map[string]interface{}))
			setting.ExtensionSetting.Extensions = append(setting.ExtensionSetting.Extensions, item)
		}
	default:
		err = fmt.Errorf("unknown ExtensionType type %#v", setting.ExtensionType)
	}
	return
}
