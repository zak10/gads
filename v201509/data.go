package v201509

type DataService struct {
	Auth
}

func NewDataService(auth *Auth) *DataService {
	return &DataService{Auth: *auth}
}
