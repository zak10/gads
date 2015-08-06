package v201409

type ConversionTrackerService struct {
	Auth
}

func NewConversionTrackerService(auth *Auth) *ConversionTrackerService {
	return &ConversionTrackerService{Auth: *auth}
}
