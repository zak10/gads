package v201409

type MutateJobService struct {
	Auth
}

func NewMutateJobService(auth *Auth) *MutateJobService {
	return &MutateJobService{Auth: *auth}
}
