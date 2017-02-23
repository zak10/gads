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

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.FeedItemScheduling
// Represents a collection of FeedItem schedules specifying all time intervals for which the feed item may serve.
// Any time range not covered by the specified FeedItemSchedules will prevent the feed item from serving during those times.
type FeedItemScheduling struct {
	feedItemSchedules int64 `xml:"https://adwords.google.com/api/adwords/cm/v201609 feedItemSchedules,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.FeedItemSchedule
// Represents a FeedItem schedule, which specifies a time interval on a given day when the feed item may serve.
// The FeedItemSchedule times are in the account's time zone.
type FeedItemSchedule struct {
	DayOfWeek   DayOfWeek    `xml:"https://adwords.google.com/api/adwords/cm/v201609 dayOfWeek,omitempty"`
	StartHour   int          `xml:"https://adwords.google.com/api/adwords/cm/v201609 startHour,omitempty"`
	StartMinute MinuteOfHour `xml:"https://adwords.google.com/api/adwords/cm/v201609 startMinute,omitempty"`
	EndHour     int          `xml:"https://adwords.google.com/api/adwords/cm/v201609 endHour,omitempty"`
	EndMinute   MinuteOfHour `xml:"https://adwords.google.com/api/adwords/cm/v201609 endMinute,omitempty"`
}
