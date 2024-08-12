package models

import (
	"time"
)

type Order struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	CustomerID  uint      `gorm:"index"`
	Customer    Customer  `gorm:"foreignKey:CustomerID"`
	OrderDate   time.Time `gorm:"autoCreateTime"`
	TotalAmount float64   `gorm:"type:decimal(10,2)"`
	Status      string    `gorm:"size:50"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
