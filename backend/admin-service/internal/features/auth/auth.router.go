package auth

import (
	"admin-service/internal/util"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	controller *AuthController
}

func NewAuthRouter(c *AuthController) *AuthRouter {
	return &AuthRouter{controller: c}
}

func (r *AuthRouter) SetupRoutes(rg *gin.RouterGroup, config *util.Config) {
	authGroup := rg.Group("/auth")

	// /auth/admins requires JWT (only root admin can register a new admin)
	authGroup.POST("/admins", RequireAuth(config, true), r.controller.Register)

	// /auth/login does not require JWT to access, but will return JWT if credentials are valid
	authGroup.POST("/login", RequireAuth(config, false), r.controller.Login)

	authGroup.POST("/logout", RequireAuth(config, true), r.controller.Logout)
}
