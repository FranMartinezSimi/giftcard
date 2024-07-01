package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	Name          string
	Address       string
	ContactEmail  string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	CompanyOrders []CompanyOrder
}
