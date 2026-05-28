package postgres

import (
	"gorm.io/gorm"
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

// --- CRUD methods for RefreshToken ---
func CreateRefreshToken(db *gorm.DB, token *RefreshToken) error {
	return db.Create(token).Error
}

func GetRefreshToken(db *gorm.DB, token string) (*RefreshToken, error) {
	var rt RefreshToken
	if err := db.Where("token = ?", token).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func RevokeRefreshTokens(db *gorm.DB, refreshToken string) error {
	return db.Model(&RefreshToken{}).Where("token = ?", refreshToken).Update("revoked", true).Error
}
