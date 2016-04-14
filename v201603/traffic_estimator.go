package v201603

type TrafficEstimatorService struct {
	Auth
}

func NewTrafficEstimatorService(auth *Auth) *TrafficEstimatorService {
	return &TrafficEstimatorService{Auth: *auth}
}
