package redirect

import (
	"example.com/shared/postgres"
	"gorm.io/gorm"
)

type RedirectRepository struct {
	db *gorm.DB
}

func NewRedirectRepository(db *gorm.DB) *RedirectRepository {
	return &RedirectRepository{db: db}
}

func (r *RedirectRepository) FindByCode(code string) (string, error) {
	var result struct {
		OriginalURL string
	}

	err := r.db.
		Model(&postgres.Link{}).
		Select("original_url").
		Where("code = ?", code).
		Take(&result).Error

	if err != nil {
		return "", err
	}

	return result.OriginalURL, nil

}
