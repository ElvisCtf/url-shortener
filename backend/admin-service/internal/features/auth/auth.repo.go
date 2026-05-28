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

// Admin CRUD
func (r *AuthRepository) CreateAdmin(admin *postgres.Admin) error {
	return postgres.CreateAdmin(r.Db, admin)
}

func (r *AuthRepository) GetAdminByID(id uint) (*postgres.Admin, error) {
	return postgres.GetAdminByID(r.Db, id)
}

func (r *AuthRepository) GetAdminByEmail(email string) (*postgres.Admin, error) {
	return postgres.GetAdminByEmail(r.Db, email)
}

func (r *AuthRepository) UpdateAdmin(admin *postgres.Admin) error {
	return postgres.UpdateAdmin(r.Db, admin)
}

func (r *AuthRepository) DeleteAdmin(id uint) error {
	return postgres.DeleteAdmin(r.Db, id)
}

// RefreshToken CRUD
func (r *AuthRepository) CreateRefreshToken(token *postgres.RefreshToken) error {
	return postgres.CreateRefreshToken(r.Db, token)
}

func (r *AuthRepository) GetRefreshToken(token string) (*postgres.RefreshToken, error) {
	return postgres.GetRefreshToken(r.Db, token)
}

func (r *AuthRepository) RevokeRefreshTokens(refreshToken string) error {
	return postgres.RevokeRefreshTokens(r.Db, refreshToken)
}
