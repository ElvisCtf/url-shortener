package auth

import (
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type Admin struct {
	gorm.Model

	Email        string `gorm:"type:varchar(255);uniqueIndex;not null"`
	HashPassword string `gorm:"type:varchar(255);not null"`
}
