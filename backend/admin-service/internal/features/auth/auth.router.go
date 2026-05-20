package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	controller *AuthController
}

func NewAuthRouter(c *AuthController) *AuthRouter {
	return &AuthRouter{controller: c}
}

func (r *AuthRouter) SetupRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")

	authGroup.POST("/register", r.controller.Register)
}
