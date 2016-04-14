package v201603

type TargetIdeaService struct {
	Auth
}

func NewTargetIdeaService(auth *Auth) *TargetIdeaService {
	return &TargetIdeaService{Auth: *auth}
}
