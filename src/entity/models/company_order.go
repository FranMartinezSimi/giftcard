package models

import (
	"time"

	"gorm.io/gorm"
)

type CompanyOrder struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CompanyID uint
	OrderID   uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Company   Company
	Order     Order
}
