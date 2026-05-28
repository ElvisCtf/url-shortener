package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/shared"
)

type AuthController struct {
	service *AuthService
	config  *shared.Config
}

func NewAuthController(s *AuthService, config *shared.Config) *AuthController {
	return &AuthController{service: s, config: config}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Register: invalid request", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.service.Register(req.Email, req.Password)
	if err != nil {
		slog.Error("Register: failed to register", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Login: invalid request", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := c.service.Login(req.Email, req.Password)
	if err != nil {
		slog.Error("Login: invalid credentials", "error", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.SetCookie("refreshToken", tokens.RefreshToken, c.config.RefreshTokenExpire, "/", c.config.Domain, shared.IsSecure(), true)
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": tokens.AccessToken,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil || refreshToken == "" {
		slog.Error("Logout: missing refresh token cookie", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token cookie"})
		return
	}

	if err := c.service.Logout(refreshToken); err != nil {
		slog.Error("Logout: failed to logout", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	ctx.SetCookie("refreshToken", "", -1, "/", c.config.Domain, shared.IsSecure(), true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil || refreshToken == "" {
		slog.Error("Refresh: missing refresh token cookie", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token cookie"})
		return
	}

	tokens, err := c.service.Refresh(refreshToken)
	if err != nil {
		slog.Error("Refresh: failed to refresh token", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	ctx.SetCookie("refreshToken", tokens.RefreshToken, c.config.RefreshTokenExpire, "/", c.config.Domain, shared.IsSecure(), true)
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": tokens.AccessToken,
	})
}
