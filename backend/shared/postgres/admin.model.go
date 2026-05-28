package postgres

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
