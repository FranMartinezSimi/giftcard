package app

import "errors"

var (
	ErrCampaignNotFound  = errors.New("campaign not found")
	ErrGiftCardNotFound  = errors.New("gift card not found")
	ErrGiftCardNotActive = errors.New("gift card is not active")
	ErrGiftCardExpired   = errors.New("gift card has expired")
	ErrInsufficientBalance = errors.New("insufficient gift card balance")
)
