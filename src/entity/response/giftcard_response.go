package response

type GetAllGiftCardResponse struct {
	ID             uint    `json:"id"`
	Type           string  `json:"type"`
	Balance        float64 `json:"balance"`
	ExpirationDate string  `json:"expiration_date"`
	Status         string  `json:"status"`
	IsPromotional  bool    `json:"is_promotional"`
}
