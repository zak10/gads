package v201506

type ExperimentService struct {
	Auth
}

func NewExperimentService(auth *Auth) *ExperimentService {
	return &ExperimentService{Auth: *auth}
}
