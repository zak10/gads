package v201605

type CampaignAdExtensionService struct {
	Auth
}

func NewCampaignAdExtensionService(auth *Auth) *CampaignAdExtensionService {
	return &CampaignAdExtensionService{Auth: *auth}
}
