package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(s *AuthService) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.service.Register(req.Email, req.Password)

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
