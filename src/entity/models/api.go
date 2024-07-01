package models

import (
	"time"

	"gorm.io/gorm"
)

type API struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Endpoint    string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
