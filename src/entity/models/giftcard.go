package models

import (
	"time"

	"gorm.io/gorm"
)

type GiftCard struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Type           string
	Balance        float64
	ExpirationDate time.Time
	Status         string
	IsPromotional  bool `gorm:"default:false"`
	CampaignID     uint
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Campaign       Campaign
	Transactions   []Transaction
	Inventory      []Inventory
}
