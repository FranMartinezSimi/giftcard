package models

import (
	"time"
)

type GiftCard struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	GiftCardNumber string    `gorm:"size:50;unique"`
	Type           string    `gorm:"size:50"`
	Balance        float64   `gorm:"type:decimal(10,2)"`
	ExpirationDate time.Time `gorm:"type:date"`
	Status         string    `gorm:"size:50"`
	IsPromotional  bool      `gorm:"default:false"`
	CampaignID     uint      `gorm:"index"`
	Campaign       Campaign  `gorm:"foreignKey:CampaignID"`
	Inventory      Inventory `gorm:"foreignKey:GiftCardID"`
	InventoryID    uint      `gorm:"index"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	Code           string    `gorm:"size:50;unique;not null"`
	InitialBalance float64   `gorm:"type:decimal(10,2)"`
	ActivationDate time.Time `gorm:"type:timestamp"`
	LastUsedDate   time.Time `gorm:"type:timestamp"`
	PinCode        string    `gorm:"size:6"`
	MaxUses        int       `gorm:"type:int"`
	CurrentUses    int       `gorm:"type:int;default:0"`
}
