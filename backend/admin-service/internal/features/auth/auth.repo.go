package auth

import (
	"admin-service/internal/util/database"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{Db: db}
}

func (r *AuthRepository) Create(admin *database.Admin) error {
	return r.Db.Create(admin).Error
}

func (r *AuthRepository) GetByID(id uint) (*database.Admin, error) {
	var admin database.Admin
	if err := r.Db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AuthRepository) GetByEmail(email string) (*database.Admin, error) {
	var admin database.Admin
	if err := r.Db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AuthRepository) Update(admin *database.Admin) error {
	return r.Db.Save(admin).Error
}

func (r *AuthRepository) Delete(id uint) error {
	return r.Db.Delete(&database.Admin{}, id).Error
}

func (r *AuthRepository) SaveRefreshToken(token *database.RefreshToken) error {
	return r.Db.Create(token).Error
}
