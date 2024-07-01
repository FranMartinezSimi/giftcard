package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	GiftCardID      uint
	Amount          float64
	TransactionType string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	GiftCard        GiftCard
}
