package repository

import "gorm.io/gorm"

type GiftCardRepository struct {
	gorm *gorm.DB
}

func NewGiftCardRepository(gorm *gorm.DB) *GiftCardRepository {
	return &GiftCardRepository{gorm: gorm}
}

func (c *GiftCardRepository) CreateGiftCard() {
}
