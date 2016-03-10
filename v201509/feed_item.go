package v201509

type FeedItemService struct {
	Auth
}

func NewFeedItemService(auth *Auth) *FeedItemService {
	return &FeedItemService{Auth: *auth}
}
