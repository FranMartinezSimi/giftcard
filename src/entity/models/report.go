package models

import (
	"time"
)

type Report struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:255"`
	Filters     string    `gorm:"type:text"`
	GeneratedAt time.Time `gorm:"autoCreateTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
