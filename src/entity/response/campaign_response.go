package response

type CampaignResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	DiscountPercentage string `json:"discount_percentage"`
}
