package models

import (
	"time"
)

type Transaction struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	GiftCardID      uint      `gorm:"index"`
	GiftCard        GiftCard  `gorm:"foreignKey:GiftCardID"`
	Amount          float64   `gorm:"type:decimal(10,2)"`
	TransactionType string    `gorm:"size:50"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
