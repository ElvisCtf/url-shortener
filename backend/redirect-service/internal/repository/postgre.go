package repository

import (
	"fmt"
	"log"

	"redirect-service/internal/model"
	"redirect-service/internal/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreRepo struct {
	db *gorm.DB
}

func newPostgreRepo() *PostgreRepo {
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

	return &PostgreRepo{db: db}
}


func (repo *PostgreRepo) FindByCode(code string) (string, error) {
    var result struct {
        OriginalURL string
    }

    err := repo.db.
        Model(&model.Link{}).
        Select("original_url").
        Where("code = ?", code).
        Take(&result).Error

    if err != nil {
		log.Println("MONKE 1")
        return "", err
    }

    return result.OriginalURL, nil
}
