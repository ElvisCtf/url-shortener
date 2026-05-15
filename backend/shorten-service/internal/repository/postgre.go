package repository

import (
	"fmt"
	"log"

	"shorten-service/internal/model"
	"shorten-service/internal/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

type PostgreRepo struct {
    db *gorm.DB
}

func NewPostgreRepo() *PostgreRepo {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		util.Env("DB_HOST", "postgres16"),
		util.Env("DB_USER", "postgres16"),
		util.Env("DB_PASSWORD", "postgres16"),
		util.Env("DB_NAME", "urlshortener"),
		util.Env("DB_PORT", "5432"),
		util.Env("DB_SSLMODE", "disable"),
		util.Env("TZ", "UTC"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil { 
		log.Fatal(err)
		return nil
	}

	if migrateErr := db.AutoMigrate(&model.Link{}); migrateErr != nil { 
		log.Fatal(err)
	}

	return &PostgreRepo{db: db}
}

func (repo *PostgreRepo) Save(originalURL string) (string, error) {
    link := &model.Link{OriginalURL: originalURL}

    // use UPSERT to insert new record
	// if new URL, then insert and return the code
	// if old URL, then do a no-op update and return the code 
    err := repo.db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "original_url"}},
			DoUpdates: clause.Assignments(map[string]any{
					"original_url": gorm.Expr("EXCLUDED.original_url"),
			}),
		}).Create(link).Error
		
	if err != nil {
		return "", err
    }

    return link.Code, nil
}