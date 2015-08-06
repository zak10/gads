package v201409

type FeedMappingService struct {
	Auth
}

func NewFeedMappingService(auth *Auth) *FeedMappingService {
	return &FeedMappingService{Auth: *auth}
}
