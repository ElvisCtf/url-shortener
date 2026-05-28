package postgres

import (
	"time"
)

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	AdminID   uint      `gorm:"not null;index"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null;index"`
	CreatedAt time.Time
	Revoked   bool `gorm:"default:false;not null;index"`
}
