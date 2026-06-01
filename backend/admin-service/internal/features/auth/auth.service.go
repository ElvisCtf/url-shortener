package auth

import (
	"errors"
	"time"

	"example.com/shared"
	"example.com/shared/postgres"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   *AuthRepository
	config *shared.Config
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// typed JWT claims struct used for both access and refresh tokens.
type AdminClaims struct {
	AdminID uint   `json:"admin_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthService(repo *AuthRepository, config *shared.Config) *AuthService {
	return &AuthService{repo: repo, config: config}
}

func (s *AuthService) Register(email, password string) (bool, error) {
	hashPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return false, hashErr
	}

	if dbErr := s.repo.CreateAdmin(&postgres.Admin{
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

func (s *AuthService) getAdminByRefreshToken(refreshToken string) (*postgres.Admin, error) {
	// Step 1: Parse and verify JWT with typed claims
	claims := &AdminClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token (JWT parse failed)")
	}

	adminID := claims.AdminID

	// Step 2: Search DB for refresh token record
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

func (s *AuthService) generateJWTs(admin *postgres.Admin) (*TokenPair, error) {
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
	rt := &postgres.RefreshToken{
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
	claims := AdminClaims{
		AdminID: adminID,
		Email:   email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
