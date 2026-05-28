package auth

import (
	"example.com/shared"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	controller *AuthController
}

func NewAuthRouter(c *AuthController) *AuthRouter {
	return &AuthRouter{controller: c}
}

func (r *AuthRouter) SetupRoutes(rg *gin.RouterGroup, config *shared.Config) {
	authGroup := rg.Group("/auth")

	// /auth/admins requires JWT (only root admin can register a new admin)
	authGroup.POST("/admins", RequireAccessToken(config, true), r.controller.Register)

	authGroup.POST("/login", RequireAccessToken(config, false), r.controller.Login)

	authGroup.POST("/logout", RequireAccessToken(config, true), r.controller.Logout)

	authGroup.POST("/refresh", RequireAccessToken(config, false), r.controller.Refresh)
}
