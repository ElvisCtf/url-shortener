package postgres

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"time"

	"gorm.io/gorm"

	"example.com/shared/util"
)

type Link struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"size:16;uniqueIndex;not null"`
	OriginalURL string `gorm:"uniqueIndex;not null"`
	Clicks      uint   `gorm:"default:0;not null"`
	Active      bool   `gorm:"default:true;not null"`
	CreatedAt   time.Time
}

// BeforeCreate generates a random base62 code before the INSERT operation and checks for uniqueness to avoid collisions.
// Retries up to 5 times in the unlikely event of a code collision.
func (l *Link) BeforeCreate(tx *gorm.DB) error {
	if l.Code != "" {
		return nil
	}
	for range 5 {
		var b [8]byte
		if _, err := rand.Read(b[:]); err != nil {
			return err
		}
		code := util.EncodeBase62(uint(binary.BigEndian.Uint64(b[:])))

		var count int64
		if err := tx.Model(&Link{}).Where("code = ?", code).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			l.Code = code
			return nil
		}
	}
	return errors.New("failed to generate a unique link code after 5 attempts")
}
