package shorten

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"example.com/shared/postgres"
)

type ShortenRepository struct {
	Db *gorm.DB
}

func NewShortenRepository(db *gorm.DB) *ShortenRepository {
	return &ShortenRepository{Db: db}
}

func (r *ShortenRepository) Save(originalURL string) (string, error) {
	link := &postgres.Link{OriginalURL: originalURL}

	// use UPSERT to insert new record
	// if new URL, then insert and return the code
	// if old URL, then do a no-op update and return the code
	err := r.Db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "original_url"}},
		DoUpdates: clause.Assignments(map[string]any{
			"original_url": gorm.Expr("EXCLUDED.original_url"),
		}),
	}).Create(link).Error

	if err != nil {
		return "", err
	}

	return link.Code, nil

}
