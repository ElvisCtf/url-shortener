package database

import (
	"example.com/admin-service/internal/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDB(cfg *util.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.TZ,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate Admin and RefreshToken models
	if err := db.AutoMigrate(&Admin{}, &RefreshToken{}); err != nil {
		return nil, err
	}

	return db, nil
}
