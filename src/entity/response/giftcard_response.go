package response

type GetAllGiftCardResponse struct {
	ID             uint    `json:"id"`
	GiftCardNumber string  `json:"gift_card_number"`
	Type           string  `json:"type"`
	Balance        float64 `json:"balance"`
	ExpirationDate string  `json:"expiration_date"`
	Status         string  `json:"status"`
	IsPromotional  bool    `json:"is_promotional"`
}

type UseGiftCardAmountResponse struct {
	GiftCardNumber string  `json:"gift_card_number"`
	Balance        float64 `json:"balance"`
	IsUsed         bool    `json:"is_used"` // Indicates if the amount was successfully used
	Message        string  `json:"message,omitempty"` // Optional message, e.g., for errors or status
}
