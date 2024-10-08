package models

import (
	"time"
)

type GiftCard struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Type           string    `gorm:"size:50"`
	Balance        float64   `gorm:"type:decimal(10,2)"`
	ExpirationDate time.Time `gorm:"type:date"`
	Status         string    `gorm:"size:50"`
	IsPromotional  bool      `gorm:"default:false"`
	CampaignID     uint      `gorm:"index"`
	Campaign       Campaign  `gorm:"foreignKey:CampaignID"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
