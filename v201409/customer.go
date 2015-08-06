package v201409

type CustomerService struct {
	Auth
}

func NewCustomerService(auth *Auth) *CustomerService {
	return &CustomerService{Auth: *auth}
}
