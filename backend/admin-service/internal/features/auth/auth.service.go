package auth

import (
	"admin-service/internal/util/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *AdminRepository
}

func NewAuthService(repo *AdminRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(email, password string) (bool, error) {
	hashPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return false, hashErr
	}

	if dbErr := s.repo.Create(&database.Admin{
		Email:        email,
		HashPassword: string(hashPassword),
	}); dbErr != nil {
		return false, dbErr
	}

	return true, nil
}
