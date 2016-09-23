package v201607

type ExperimentService struct {
	Auth
}

func NewExperimentService(auth *Auth) *ExperimentService {
	return &ExperimentService{Auth: *auth}
}
