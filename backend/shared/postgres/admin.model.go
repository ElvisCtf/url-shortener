package postgres

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"type:varchar(255);uniqueIndex;not null"`
	HashPassword string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	RefreshTokens []RefreshToken `gorm:"foreignKey:AdminID;constraint:OnDelete:CASCADE"`
}

// --- CRUD methods for Admin ---
func CreateAdmin(db *gorm.DB, admin *Admin) error {
	return db.Create(admin).Error
}

func GetAdminByID(db *gorm.DB, id uint) (*Admin, error) {
	var admin Admin
	if err := db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func GetAdminByEmail(db *gorm.DB, email string) (*Admin, error) {
	var admin Admin
	if err := db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func UpdateAdmin(db *gorm.DB, admin *Admin) error {
	return db.Save(admin).Error
}

func DeleteAdmin(db *gorm.DB, id uint) error {
	return db.Delete(&Admin{}, id).Error
}
