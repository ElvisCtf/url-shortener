package database

import (
	"time"
)

type Admin struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"type:varchar(255);uniqueIndex;not null"`
	HashPassword string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	RefreshTokens []RefreshToken `gorm:"foreignKey:AdminID;constraint:OnDelete:CASCADE"`
}

type RefreshToken struct {
	ID        string    `gorm:"primaryKey;autoIncrement"`
	AdminID   uint      `gorm:"not null;index"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null;index"`
	CreatedAt time.Time
	Revoked   bool `gorm:"default:false;not null;index"`
}
