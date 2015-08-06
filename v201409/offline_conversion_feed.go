package v201409

type OfflineConversionService struct {
	Auth
}

func NewOfflineConversionService(auth *Auth) *OfflineConversionService {
	return &OfflineConversionService{Auth: *auth}
}
