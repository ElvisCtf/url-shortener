package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"example.com/shared"
	"example.com/shared/postgres"
)

func main() {
	config, err := shared.LoadConfig()
	config.DBHost = "localhost"
	config.DBPort = "5430"

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	db, err := postgres.InitGormDB(config)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	if err := ResetDatabase(db); err != nil {
		log.Fatalf("failed to reset database: %v", err)
	}
	fmt.Println("Database reset successfully!")
}

// ResetDatabase drops all tables and re-applies migrations for Admin, RefreshToken, and Link models.
func ResetDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&postgres.Admin{}, &postgres.RefreshToken{}, &postgres.Link{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&postgres.Admin{}, &postgres.RefreshToken{}, &postgres.Link{}); err != nil {
		return err
	}
	return nil
}
