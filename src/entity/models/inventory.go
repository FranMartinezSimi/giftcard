package models

import (
	"time"
)

type Inventory struct {
	ID           uint     `gorm:"primaryKey;autoIncrement"`
	GiftCardID   uint     `gorm:"index"`
	GiftCard     GiftCard `gorm:"foreignKey:GiftCardID"`
	LocationType string   `gorm:"size:50"`
	BinLocation  string   `gorm:"size:50"`
	Quantity     int
	Status       string    `gorm:"size:50"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
