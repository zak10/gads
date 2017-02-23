package v201609

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
type AdGroupExtensionSettingService struct {
	Auth
}

func NewAdGroupExtensionSettingService(auth *Auth) *AdGroupExtensionSettingService {
	return &AdGroupExtensionSettingService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.AdGroupExtensionSettingPage
// Contains a subset of AdGroupExtensionSetting objects resulting from a AdGroupExtensionSettingService#get call.
type AdGroupExtensionSettingPage struct {
	Page
	Entries []AdGroupExtensionSetting `xml:"https://adwords.google.com/api/adwords/cm/v201609 entries,omitempty"`
}
type AdGroupExtensionSettingPage4CallFeed struct {
	Page
	Entries []AdGroupExtensionSetting4CallFeed `xml:"https://adwords.google.com/api/adwords/cm/v201609 entries,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.AdGroupExtensionSetting
// An AdGroupExtensionSetting is used to add or modify extensions being served for the specified ad group.
type AdGroupExtensionSetting struct {
	AdGroupId        int64            `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupId,omitempty"`
	ExtensionType    FeedType         `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting ExtensionSetting `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}
type AdGroupExtensionSetting4CallFeed struct {
	AdGroupId        int64                     `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupId,omitempty"`
	ExtensionType    FeedType                  `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting ExtensionSetting4CallFeed `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}
