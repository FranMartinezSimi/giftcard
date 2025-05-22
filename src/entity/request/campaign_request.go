package request

import "time"

type CreateCampaignRequest struct {
	Name               string    `json:"name" validate:"required,min=3,max=255"`
	Description        string    `json:"description" validate:"max=1000"`
	StartDate          time.Time `json:"start_date" validate:"required"`
	EndDate            time.Time `json:"end_date" validate:"required,gtfield=StartDate"`
	IsEnabled          bool      `json:"is_enabled"`
	DiscountPercentage float64   `json:"discount_percentage" validate:"required,min=0,max=100"`
}

type UpdateCampaignRequest struct {
	Name               string    `json:"name" validate:"required,min=3,max=255"`
	Description        string    `json:"description" validate:"max=1000"`
	StartDate          time.Time `json:"start_date" validate:"required"`
	EndDate            time.Time `json:"end_date" validate:"required,gtfield=StartDate"`
	IsEnabled          bool      `json:"is_enabled"`
	DiscountPercentage float64   `json:"discount_percentage" validate:"required,min=0,max=100"`
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
