package auth

import (
	"admin-service/internal/util"
	"admin-service/internal/util/database"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   *AuthRepository
	config *util.Config
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func NewAuthService(repo *AuthRepository, config *util.Config) *AuthService {
	return &AuthService{repo: repo, config: config}
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

func (s *AuthService) Login(email, password string) (*TokenPair, error) {
	admin, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.HashPassword), []byte(password)); err != nil {
		return nil, err
	}

	accessTokenExpire := s.config.AccessTokenExpire
	refreshTokenExpire := s.config.RefreshTokenExpire

	accessToken, err := generateJWTWithSecret(admin.ID, admin.Email, accessTokenExpire, s.config.JWTSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateJWTWithSecret(admin.ID, admin.Email, refreshTokenExpire, s.config.JWTSecret)
	if err != nil {
		return nil, err
	}

	// Save refresh token in DB
	rt := &database.RefreshToken{
		AdminID:   admin.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Duration(refreshTokenExpire) * time.Second),
		CreatedAt: time.Now(),
		Revoked:   false,
	}
	if err := s.repo.SaveRefreshToken(rt); err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateJWTWithSecret(adminID uint, email string, seconds int, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}
	duration := time.Duration(seconds) * time.Second
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"email":    email,
		"exp": jwt.NewNumericDate(time.Now().Add(duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
