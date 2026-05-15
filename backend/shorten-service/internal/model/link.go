package model

import (
    "time"

    "shorten-service/internal/util"

    "gorm.io/gorm"
)

type Link struct {
    ID          uint64    `gorm:"primaryKey;autoIncrement"`
    Code        string    `gorm:"size:16;uniqueIndex;not null"`
    OriginalURL string    `gorm:"not null;uniqueIndex"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (l *Link) AfterCreate(tx *gorm.DB) (err error) {
    if l.Code != "" {
        return nil
    }
    code := util.EncodeBase62(l.ID)
    l.Code = code
    return tx.Model(l).Update("code", code).Error
}
