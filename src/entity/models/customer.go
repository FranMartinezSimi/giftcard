package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Orders    []Order
}
