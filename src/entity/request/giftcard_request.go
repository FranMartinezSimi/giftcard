package request

type CreateGiftCardRequest struct {
	Type           string  `json:"type"`
	Balance        float64 `json:"balance"`
	ExpirationDate string  `json:"expiration_date"`
	Status         string  `json:"status"`
	IsPromotional  bool    `json:"is_promotional"`
	CampaignID     uint    `json:"campaign_id"`
}

type UpdateGiftCardRequest struct {
	Type           string  `json:"type"`
	Balance        float64 `json:"balance"`
	ExpirationDate string  `json:"expiration_date"`
	Status         string  `json:"status"`
	IsPromotional  bool    `json:"is_promotional"`
}
