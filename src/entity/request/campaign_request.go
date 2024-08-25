package request

import "time"

type CreateCampaignRequest struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	IsEnabled          bool      `json:"is_enabled"`
	DiscountPercentage float64   `json:"discount_percentage"`
}

type UpdateCampaignRequest struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	IsEnabled          bool      `json:"is_enabled"`
	DiscountPercentage float64   `json:"discount_percentage"`
}

type FullTextSearchCampaignRequest struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	IsEnabled          bool      `json:"is_enabled"`
	DiscountPercentage float64   `json:"discount_percentage"`
}
