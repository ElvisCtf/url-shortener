package redirect

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"example.com/shared/postgres"
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

func (r *RedirectRepository) AddClicks(ctx context.Context, code string) error {
	return r.db.WithContext(ctx).Model(&postgres.Link{}).
		Where("code = ?", code).
		Update("clicks", gorm.Expr("clicks + ?", 1)).Error
}
