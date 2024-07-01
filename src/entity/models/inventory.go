package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	GiftCardID   uint `gorm:"foreignKey:ID"`
	LocationType string
	BinLocation  string
	Quantity     int
	Status       string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	GiftCard     GiftCard
}
