package models

import (
	"time"
)

type Campaign struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	CampaignUUID       string    `gorm:"size:255"`
	Name               string    `gorm:"size:255"`
	Description        string    `gorm:"type:text"`
	StartDate          time.Time `gorm:"type:date"`
	EndDate            time.Time `gorm:"type:date"`
	IsEnabled          bool      `gorm:"default:true"`
	DiscountPercentage float64   `gorm:"type:decimal(5,2)"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
