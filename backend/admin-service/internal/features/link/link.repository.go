package link

import (
	"example.com/shared/postgres"
	"gorm.io/gorm"
)

type LinkRepository struct {
	Db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{Db: db}
}

func (r *LinkRepository) GetPaginated(page, pageSize int, order string) ([]postgres.Link, int64, error) {
	var links []postgres.Link
	var totalCount int64

	// Count total records for metadata
	if err := r.Db.Model(&postgres.Link{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply default values if missing
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if order == "" {
		order = "id desc"
	}

	offset := (page - 1) * pageSize

	err := r.Db.Order(order).Limit(pageSize).Offset(offset).Find(&links).Error
	return links, totalCount, err
}
