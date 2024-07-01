package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Filters   string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
