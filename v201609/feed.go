package v201609

type FeedService struct {
	Auth
}

func NewFeedService(auth *Auth) *FeedService {
	return &FeedService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.Feed.Type
// Feed hard type. Values coincide with placeholder type id.
// Enum: NONE, SITELINK, CALL, APP, REVIEW, AD_CUSTOMIZER, CALLOUT, STRUCTURED_SNIPPET, PRICE
type FeedType string
