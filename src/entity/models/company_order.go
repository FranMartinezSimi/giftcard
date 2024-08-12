package models

import (
	"time"
)

type CompanyOrder struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CompanyID uint      `gorm:"index"`
	Company   Company   `gorm:"foreignKey:CompanyID"`
	OrderID   uint      `gorm:"index"`
	Order     Order     `gorm:"foreignKey:OrderID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
