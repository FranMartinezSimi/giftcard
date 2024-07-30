package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model

	ID                 uint `gorm:"primaryKey"`
	Name               string
	Description        string
	StartDate          time.Time
	EndDate            time.Time
	DiscountPercentage float64
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
	GiftCards          []GiftCard
}

// Error implements error.
func (c *Campaign) Error() string {
	panic("unimplemented")
}
