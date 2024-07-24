package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	GiftCards    []GiftCard
	LocationType string
	BinLocation  string
	Quantity     int
	Status       string
	CampaignID   uint `gorm:"foreignKey:Campaign.ID"`
	Campaign     Campaign
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
