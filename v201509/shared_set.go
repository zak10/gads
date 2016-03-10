package v201509

type SharedSetService struct {
	Auth
}

func NewSharedSetService(auth *Auth) *SharedSetService {
	return &SharedSetService{Auth: *auth}
}
