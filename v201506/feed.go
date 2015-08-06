package v201506

type FeedService struct {
	Auth
}

func NewFeedService(auth *Auth) *FeedService {
	return &FeedService{Auth: *auth}
}
