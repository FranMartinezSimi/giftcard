package models

import (
	"time"
)

type Customer struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	Phone     string    `gorm:"size:20"`
	Address   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
