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
	group := rg.Group("api/v1/auth")

	// /auth/admins requires JWT (only root admin can register a new admin)
	group.POST("/admins", RequireAccessToken(config), r.controller.Register)
	group.POST("/logout", RequireAccessToken(config), r.controller.Logout)

	group.POST("/login", r.controller.Login)
	group.POST("/refresh", r.controller.Refresh)
}
