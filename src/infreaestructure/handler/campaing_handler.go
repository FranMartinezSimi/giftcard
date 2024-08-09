package handler

import (
	"GiftWize/src/app"
)

type CampaignHandler struct {
	useCase app.Campaign
}

func NewCampaignHandler(useCase app.Campaign) *CampaignHandler {
	return &CampaignHandler{
		useCase: useCase,
	}
}

func (h *CampaignHandler) CreateCampaign() {
	// Implementación de la función
}
