package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"admin-service/internal/features/auth"
)

func InitGormDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.TZ,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate Admin model
	if err := db.AutoMigrate(&auth.Admin{}); err != nil {
		return nil, err
	}
	return db, nil
}
