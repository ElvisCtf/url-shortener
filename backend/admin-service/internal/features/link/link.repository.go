package link

import (
	"gorm.io/gorm"

	"example.com/shared/postgres"
)

type LinkRepository struct {
	Db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{Db: db}
}

var allowedOrders = map[string]string{
	"asc":  "id ASC",
	"desc": "id DESC",
}

func (r *LinkRepository) GetPaginated(page, pageSize int, order string) ([]postgres.Link, int64, error) {
	var links []postgres.Link
	var totalCount int64

	// Count total records for metadata
	if err := r.Db.Model(&postgres.Link{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Whitelist order values to prevent SQL injection
	safeOrder, ok := allowedOrders[order]
	if !ok {
		safeOrder = "id DESC"
	}

	offset := (page - 1) * pageSize

	err := r.Db.Order(safeOrder).Limit(pageSize).Offset(offset).Find(&links).Error
	return links, totalCount, err
}

func (r *LinkRepository) BulkUpdateActive(ids []uint, active bool) error {
	return r.Db.Model(&postgres.Link{}).Where("id IN ?", ids).Update("active", active).Error
}
