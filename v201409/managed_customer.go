package v201409

type ManagedCustomerService struct {
	Auth
}

func NewManagedCustomerService(auth *Auth) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth}
}
