package postgres

import (
	"time"

	"gorm.io/gorm"

	"example.com/shared/util"
)

type Link struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"size:16;uniqueIndex;not null"`
	OriginalURL string `gorm:"uniqueIndex;not null"`
	CreatedAt   time.Time
}

func (l *Link) AfterCreate(tx *gorm.DB) (err error) {
	if l.Code != "" {
		return nil
	}
	code := util.EncodeBase62(l.ID)
	l.Code = code
	return tx.Model(l).Update("code", code).Error
}

func GetPaginatedLinks(db *gorm.DB, page, pageSize int, order string) ([]*Link, error) {
	var links []*Link
	offset := (page - 1) * pageSize
	orderStr := "CreatedAt DESC"
	if order == "asc" {
		orderStr = "CreatedAt ASC"
	}
	if err := db.Limit(pageSize).Offset(offset).Order(orderStr).Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}
