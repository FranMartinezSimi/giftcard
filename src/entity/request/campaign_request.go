package request

import "time"

type CreateCampaignRequest struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	DiscountPercentage float64   `json:"discount_percentage"`
}
