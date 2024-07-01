package models

import (
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Value     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
