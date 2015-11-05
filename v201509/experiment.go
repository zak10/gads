package v201509

type ExperimentService struct {
	Auth
}

func NewExperimentService(auth *Auth) *ExperimentService {
	return &ExperimentService{Auth: *auth}
}
