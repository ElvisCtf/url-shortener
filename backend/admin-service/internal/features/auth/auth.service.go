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

	if dbErr := s.repo.CreateAdmin(&database.Admin{
		Email:        email,
		HashPassword: string(hashPassword),
	}); dbErr != nil {
		return false, dbErr
	}

	return true, nil
}

func (s *AuthService) Login(email, password string) (*TokenPair, error) {
	admin, err := s.repo.GetAdminByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.HashPassword), []byte(password)); err != nil {
		return nil, err
	}

	return s.generateJWTs(admin)
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.repo.RevokeRefreshTokens(refreshToken)
}

func (s *AuthService) Refresh(refreshToken string) (*TokenPair, error) {
	admin, err := s.getAdminByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	err = s.repo.RevokeRefreshTokens(refreshToken)
	if err != nil {
		return nil, err
	}

	return s.generateJWTs(admin)
}

func (s *AuthService) getAdminByRefreshToken(refreshToken string) (*database.Admin, error) {
	// Step 1: Parse and verify JWT
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token (JWT parse failed)")
	}

	// Step 2: Extract admin_id from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	adminIDFloat, ok := claims["admin_id"].(float64)
	if !ok {
		return nil, errors.New("admin_id not found in token claims")
	}
	adminID := uint(adminIDFloat)

	// Step 3: Search DB for refresh token record
	rt, err := s.repo.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("refresh token not found in DB")
	}
	if rt.Revoked {
		return nil, errors.New("refresh token is revoked")
	}
	if time.Now().After(rt.ExpiresAt) {
		return nil, errors.New("refresh token is expired")
	}

	// Step 4: Cross-check admin_id
	if rt.AdminID != adminID {
		return nil, errors.New("admin_id mismatch between token and DB")
	}

	// Fetch admin from DB (optional, but usually needed for login/rotation)
	admin, err := s.repo.GetAdminByID(adminID)
	if err != nil {
		return nil, errors.New("admin not found")
	}
	return admin, nil
}

func (s *AuthService) generateJWTs(admin *database.Admin) (*TokenPair, error) {
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
	if err := s.repo.CreateRefreshToken(rt); err != nil {
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
		"exp":      jwt.NewNumericDate(time.Now().Add(duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
