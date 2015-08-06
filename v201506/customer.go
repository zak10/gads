package v201506

type CustomerService struct {
	Auth
}

func NewCustomerService(auth *Auth) *CustomerService {
	return &CustomerService{Auth: *auth}
}
