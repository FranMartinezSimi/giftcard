package request

type CreateGiftCardRequest struct {
	// Consider using oneof for predefined types: e.g., "virtual", "physical"
	Type           string  `json:"type" validate:"required"`
	Balance        float64 `json:"balance" validate:"required,min=0"`
	// Add custom validation for future date if needed
	ExpirationDate string  `json:"expiration_date" validate:"required"`
	// Consider using oneof for predefined statuses: e.g., "active", "inactive", "expired"
	Status         string  `json:"status" validate:"required"`
	IsPromotional  bool    `json:"is_promotional"`
	CampaignID     uint    `json:"campaign_id" validate:"omitempty,gt=0"`
}

type UpdateGiftCardRequest struct {
	// Consider using oneof for predefined types: e.g., "virtual", "physical"
	Type           string  `json:"type" validate:"required"`
	Balance        float64 `json:"balance" validate:"required,min=0"`
	// Add custom validation for future date if needed
	ExpirationDate string  `json:"expiration_date" validate:"required"`
	// Consider using oneof for predefined statuses: e.g., "active", "inactive", "expired"
	Status         string  `json:"status" validate:"required"`
	IsPromotional  bool    `json:"is_promotional"`
}
