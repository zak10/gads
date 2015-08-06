package v201409

type GeoLocationService struct {
	Auth
}

func NewGeoLocationService(auth *Auth) *GeoLocationService {
	return &GeoLocationService{Auth: *auth}
}
