package v201609

type FeedItemService struct {
	Auth
}

func NewFeedItemService(auth *Auth) *FeedItemService {
	return &FeedItemService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.FeedItem.Status
// ENABLED, REMOVED, UNKNOWN
type FeedItemStatus string

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.FeedItemDevicePreference
// Represents a FeedItem device preference
type FeedItemDevicePreference struct {
	DevicePreference int64 `xml:"https://adwords.google.com/api/adwords/cm/v201609 devicePreference,omitempty"`
}
