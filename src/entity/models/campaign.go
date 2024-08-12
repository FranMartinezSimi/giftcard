package models

import (
	"time"
)

type Campaign struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	Name               string    `gorm:"size:255"`
	Description        string    `gorm:"type:text"`
	StartDate          time.Time `gorm:"type:date"`
	EndDate            time.Time `gorm:"type:date"`
	DiscountPercentage float64   `gorm:"type:decimal(5,2)"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
