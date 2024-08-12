package models

import (
	"time"
)

type AuditLog struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Action    string `gorm:"type:text"`
	UserID    uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
