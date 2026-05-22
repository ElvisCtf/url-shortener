package main

import (
	"admin-service/internal/util"
	"admin-service/internal/util/database"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func main() {
	config, err := util.LoadConfig()
	config.DBHost = "localhost"
	config.DBPort = "5430"

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	db, err := database.InitGormDB(config)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	if err := ResetDatabase(db); err != nil {
		log.Fatalf("failed to reset database: %v", err)
	}
	fmt.Println("Database reset successfully!")
}

// ResetDatabase drops all tables and re-applies migrations for Admin and RefreshToken models.
func ResetDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&database.Admin{}, &database.RefreshToken{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&database.Admin{}, &database.RefreshToken{}); err != nil {
		return err
	}
	return nil
}
