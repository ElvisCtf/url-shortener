package redirect

import (
	"errors"
	"example.com/shared/postgres"
	"gorm.io/gorm"
)

var (
	ErrLinkNotFound = errors.New("link code not found")
	ErrLinkInactive = errors.New("link is inactive")
)

type RedirectRepository struct {
	db *gorm.DB
}

func NewRedirectRepository(db *gorm.DB) *RedirectRepository {
	return &RedirectRepository{db: db}
}

func (r *RedirectRepository) FindByCode(code string) (string, error) {
	var link postgres.Link

	err := r.db.Where("code = ?", code).First(&link).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrLinkNotFound
		}
		return "", err
	}

	if !link.Active {
		return "", ErrLinkInactive
	}

	return link.OriginalURL, nil
}
