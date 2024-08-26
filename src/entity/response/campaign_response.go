package response

type CampaignResponse struct {
	ID                 uint   `json:"id"`
	Uuid               string `json:"uuid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	IsEnabled          bool   `json:"is_enabled"`
	DiscountPercentage string `json:"discount_percentage"`
	CreatedAt          string `json:"created_at"`
}
