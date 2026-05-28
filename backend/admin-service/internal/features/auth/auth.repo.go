package auth

import (
	"example.com/shared/postgres"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{Db: db}
}

func (r *AuthRepository) CreateAdmin(admin *postgres.Admin) error {
	return r.Db.Create(admin).Error
}

func (r *AuthRepository) GetAdminByID(id uint) (*postgres.Admin, error) {
	var admin postgres.Admin
	if err := r.Db.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AuthRepository) GetAdminByEmail(email string) (*postgres.Admin, error) {
	var admin postgres.Admin
	if err := r.Db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AuthRepository) UpdateAdmin(admin *postgres.Admin) error {
	return r.Db.Save(admin).Error
}

func (r *AuthRepository) DeleteAdmin(id uint) error {
	return r.Db.Delete(&postgres.Admin{}, id).Error
}

func (r *AuthRepository) CreateRefreshToken(token *postgres.RefreshToken) error {
	return r.Db.Create(token).Error
}

func (r *AuthRepository) GetRefreshToken(token string) (*postgres.RefreshToken, error) {
	var rt postgres.RefreshToken
	if err := r.Db.Where("token = ?", token).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *AuthRepository) RevokeRefreshTokens(refreshToken string) error {
	return r.Db.Model(&postgres.RefreshToken{}).Where("token = ?", refreshToken).Update("revoked", true).Error
}
