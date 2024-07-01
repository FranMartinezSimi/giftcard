package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	CustomerID    uint
	OrderDate     time.Time
	TotalAmount   float64
	Status        string
	Customer      Customer
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	CompanyOrders []CompanyOrder
}
